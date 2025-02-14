package status

import (
	"context"
	"fmt"

	"github.com/redhat-developer/app-services-cli/pkg/color"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cluster"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/logging"

	"github.com/spf13/cobra"

	// Get all auth schemes
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Options struct {
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)
	IO         *iostreams.IOStreams

	kubeconfig string
}

func NewStatusCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("cluster.status.cmd.use"),
		Short:   localizer.MustLocalizeFromID("cluster.status.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("cluster.status.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("cluster.status.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runStatus(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.kubeconfig, "kubeconfig", "", "", localizer.MustLocalizeFromID("cluster.common.flag.kubeconfig.description"))

	return cmd
}

func runStatus(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	clusterConn, err := cluster.NewKubernetesClusterConnection(connection, opts.Config, logger, opts.kubeconfig, opts.IO)
	if err != nil {
		return err
	}

	var operatorStatus string
	// Add versioning in future
	isCRDInstalled, err := clusterConn.IsRhoasOperatorAvailableOnCluster(context.Background())
	if isCRDInstalled && err != nil {
		logger.Debug(err)
	}

	if isCRDInstalled {
		operatorStatus = color.Success(localizer.MustLocalizeFromID("cluster.common.operatorInstalledMessage"))
	} else {
		operatorStatus = color.Error(localizer.MustLocalizeFromID("cluster.common.operatorNotInstalledMessage"))
	}

	currentNamespace, err := clusterConn.CurrentNamespace()
	if err != nil {
		return err
	}

	fmt.Fprintln(
		opts.IO.Out,
		localizer.MustLocalize(&localizer.Config{
			MessageID: "cluster.status.statusMessage",
			TemplateData: map[string]interface{}{
				"Namespace":      color.Info(currentNamespace),
				"OperatorStatus": operatorStatus,
			},
		}),
	)

	return nil
}
