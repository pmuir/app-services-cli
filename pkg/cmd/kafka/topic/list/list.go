package list

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	topicutil "github.com/redhat-developer/app-services-cli/pkg/kafka/topic"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/flag"
	"github.com/redhat-developer/app-services-cli/pkg/connection"

	strimziadminclient "github.com/redhat-developer/app-services-cli/pkg/api/strimzi-admin/client"

	"gopkg.in/yaml.v2"

	"github.com/redhat-developer/app-services-cli/internal/config"
	"github.com/redhat-developer/app-services-cli/pkg/cmd/factory"
	"github.com/redhat-developer/app-services-cli/pkg/dump"
	"github.com/redhat-developer/app-services-cli/pkg/iostreams"
	"github.com/redhat-developer/app-services-cli/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Config     config.IConfig
	IO         *iostreams.IOStreams
	Connection factory.ConnectionFunc
	Logger     func() (logging.Logger, error)

	kafkaID string
	output  string
}

type topicRow struct {
	Name            string `json:"name,omitempty" header:"Name"`
	PartitionsCount int    `json:"partitions_count,omitempty" header:"Partitions"`
	RetentionTime   string `json:"retention.ms,omitempty" header:"Retention time (ms)"`
	RetentionSize   string `json:"retention.bytes,omitempty" header:"Retention size (bytes)"`
}

// NewListTopicCommand gets a new command for getting kafkas.
func NewListTopicCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Config:     f.Config,
		Connection: f.Connection,
		Logger:     f.Logger,
		IO:         f.IOStreams,
	}

	cmd := &cobra.Command{
		Use:     localizer.MustLocalizeFromID("kafka.topic.list.cmd.use"),
		Short:   localizer.MustLocalizeFromID("kafka.topic.list.cmd.shortDescription"),
		Long:    localizer.MustLocalizeFromID("kafka.topic.list.cmd.longDescription"),
		Example: localizer.MustLocalizeFromID("kafka.topic.list.cmd.example"),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.output != "" {
				if err := flag.ValidateOutput(opts.output); err != nil {
					return err
				}
			}

			cfg, err := opts.Config.Load()
			if err != nil {
				return err
			}

			if !cfg.HasKafka() {
				return errors.New(localizer.MustLocalizeFromID("kafka.topic.common.error.noKafkaSelected"))
			}

			opts.kafkaID = cfg.Services.Kafka.ClusterID

			return runCmd(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", localizer.MustLocalize(&localizer.Config{
		MessageID:   "kafka.topic.common.flag.output.description",
		PluralCount: 2,
	}))

	return cmd
}

func runCmd(opts *Options) error {
	conn, err := opts.Connection(connection.DefaultConfigRequireMasAuth)
	if err != nil {
		return err
	}

	logger, err := opts.Logger()
	if err != nil {
		return err
	}

	api, kafkaInstance, err := conn.API().TopicAdmin(opts.kafkaID)
	if err != nil {
		return err
	}

	a := api.GetTopicsList(context.Background())
	topicData, httpRes, topicErr := a.Execute()

	if topicErr.Error() != "" {
		if httpRes == nil {
			return topicErr
		}

		switch httpRes.StatusCode {
		case 401:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID:   "kafka.topic.common.error.unauthorized",
				PluralCount: 2,
				TemplateData: map[string]interface{}{
					"Operation": "list",
				},
			}))
		case 403:
			return errors.New(localizer.MustLocalize(&localizer.Config{
				MessageID:   "kafka.topic.common.error.forbidden",
				PluralCount: 2,
				TemplateData: map[string]interface{}{
					"Operation": "list",
				},
			}))
		case 500:
			return errors.New(localizer.MustLocalizeFromID("kafka.topic.common.error.internalServerError"))
		case 503:
			return fmt.Errorf("%v: %w", localizer.MustLocalize(&localizer.Config{
				MessageID: "kafka.topic.common.error.unableToConnectToKafka",
				TemplateData: map[string]interface{}{
					"Name": kafkaInstance.GetName(),
				},
			}), topicErr)
		default:
			return topicErr
		}
	}

	if topicData.GetCount() == 0 {
		logger.Info(localizer.MustLocalize(&localizer.Config{
			MessageID: "kafka.topic.list.log.info.noTopics",
			TemplateData: map[string]interface{}{
				"InstanceName": kafkaInstance.GetName(),
			},
		}))

		return nil
	}

	stdout := opts.IO.Out
	switch opts.output {
	case "json":
		data, _ := json.Marshal(topicData)
		_ = dump.JSON(stdout, data)
	case "yaml", "yml":
		data, _ := yaml.Marshal(topicData)
		_ = dump.YAML(stdout, data)
	default:
		topics := topicData.GetItems()
		rows := mapTopicResultsToTableFormat(topics)
		dump.Table(stdout, rows)
	}

	return err
}

func mapTopicResultsToTableFormat(topics []strimziadminclient.Topic) []topicRow {
	var rows []topicRow = []topicRow{}

	for _, t := range topics {

		row := topicRow{
			Name:            t.GetName(),
			PartitionsCount: len(t.GetPartitions()),
		}
		for _, config := range t.GetConfig() {
			unlimitedVal := "-1 (Unlimited)"

			if *config.Key == topicutil.RetentionMsKey {
				val := config.GetValue()
				if val == "-1" {
					row.RetentionTime = unlimitedVal
				} else {
					row.RetentionTime = val
				}
			}
			if *config.Key == topicutil.RetentionSizeKey {
				val := config.GetValue()
				if val == "-1" {
					row.RetentionSize = unlimitedVal
				} else {
					row.RetentionSize = val
				}
			}
		}

		rows = append(rows, row)
	}

	return rows
}
