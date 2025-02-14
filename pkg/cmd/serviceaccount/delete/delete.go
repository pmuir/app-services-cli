package delete

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/connection"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO         *iostreams.IOStreams
	Config     config.IConfig
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)

	id    string
	force bool
}

// NewDeleteCommand creates a new command to delete a service account
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("serviceAccount.delete.cmd.use"),
		Short:   localizer.MustLocalizeFromID("serviceAccount.delete.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("serviceAccount.delete.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("serviceAccount.delete.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !opts.IO.CanPrompt() && !opts.force {
				return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
					MessageID: "flag.error.requiredWhenNonInteractive",
					TemplateData: map[string]interface{}{
						"Flag": "yes",
					},
				}))
			}

			return runDelete(opts)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", localizer.MustLocalizeFromID("serviceAccount.delete.flag.id.description"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, localizer.MustLocalizeFromID("serviceAccount.delete.flag.yes.description"))

	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func runDelete(opts *Options) (err error) {
	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	api := connection.API()
	a := api.Kafka().GetServiceAccountById(context.Background(), opts.id)
	_, httpRes, apiErr := a.Execute()

	if apiErr.Error() != "" {
		if httpRes == nil {
			return apiErr
		}

		if httpRes.StatusCode == 404 {
			return fmt.Errorf(localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.common.error.notFoundError",
				TemplateData: map[string]interface{}{
					"ID": opts.id,
				},
			}))
		}
	}

	if !opts.force {
		var confirmDelete bool
		promptConfirmDelete := &survey.Confirm{
			Message: localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.delete.input.confirmDelete.message",
				TemplateData: map[string]interface{}{
					"ID": opts.id,
				},
			}),
		}

		err = survey.AskOne(promptConfirmDelete, &confirmDelete)
		if err != nil {
			return err
		}

		if !confirmDelete {
			logger.Debug(localizer.MustLocalizeFromID("serviceAccount.delete.log.debug.deleteNotConfirmed"))
			return nil
		}
	}

	return deleteServiceAccount(opts)
}

func deleteServiceAccount(opts *Options) error {
	connection, err := opts.Connection(connection.DefaultConfigSkipMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api := connection.API()

	a := api.Kafka().DeleteServiceAccount(context.Background(), opts.id)
	_, httpRes, apiErr := a.Execute()

	if apiErr.Error() != "" {
		if httpRes == nil {
			return apiErr
		}

		switch httpRes.StatusCode {
		case 403:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "serviceAccount.common.error.forbidden",
				TemplateData: map[string]interface{}{
					"Operation": "delete",
				},
			}), apiErr)
		case 500:
			return fmt.Errorf("%v: %w", localizer.MustLocalizeFromID("serviceAccount.common.error.internalServerError"), apiErr)
		default:
			return apiErr
		}
	}

	logger.Info(localizer.MustLocalizeFromID("serviceAccount.delete.log.info.deleteSuccess"))

	return nil
}
