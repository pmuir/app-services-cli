package kafka

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
	kasclient "github.com/redhat-developer/app-services-cli/pkg/api/kas/client"
)

var (
	validNameRegexp   = regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`)
	validSearchRegexp = regexp.MustCompile(`^([a-zA-Z0-9-_%]*[a-zA-Z0-9-_%])?$`)
)

// ValidateName validates the proposed name of a Kafka instance
func ValidateName(val interface{}) error {
	name, ok := val.(string)

	if !ok {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.castError",
			TemplateData: map[string]interface{}{
				"Value": val,
				"Type":  "string",
			},
		}))
	}

	if len(name) < 1 || len(name) > 32 {
		return fmt.Errorf(localizer.MustLocalizeFromID("kafka.validation.name.error.lengthError"))
	}

	matched := validNameRegexp.MatchString(name)

	if matched {
		return nil
	}

	return errors.New(localizer.MustLocalizeFromID("kafka.validation.error.invalidName"))
}

// TransformKafkaRequestListItems modifies fields fields from a list of kafka instances
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequestListItems(items []kasclient.KafkaRequest) []kasclient.KafkaRequest {
	for i := range items {
		kafka := items[i]
		kafka = *TransformKafkaRequest(&kafka)
		items[i] = kafka
	}

	return items
}

// TransformKafkaRequest modifies fields from the KafkaRequest payload object
// The main transformation is appending ":443" to the Bootstrap Server URL
func TransformKafkaRequest(kafka *kasclient.KafkaRequest) *kasclient.KafkaRequest {
	bootstrapHost := kafka.GetBootstrapServerHost()

	if bootstrapHost == "" {
		return kafka
	}

	if !strings.HasSuffix(bootstrapHost, ":443") {
		hostURL := fmt.Sprintf("%v:443", bootstrapHost)
		kafka.SetBootstrapServerHost(hostURL)
	}

	return kafka
}

// ValidateSearchInput validates the text provided to filter the Kafka instances
func ValidateSearchInput(val interface{}) error {
	search, ok := val.(string)

	if !ok {
		return errors.New(localizer.MustLocalize(&localizer.Config{
			MessageID: "common.error.castError",
			TemplateData: map[string]interface{}{
				"Value": val,
				"Type":  "string",
			},
		}))
	}

	matched := validSearchRegexp.MatchString(search)

	if matched {
		return nil
	}

	return errors.New(localizer.MustLocalize(&localizer.Config{
		MessageID: "kafka.validation.error.invalidSearchValue",
		TemplateData: map[string]interface{}{
			"Search": search,
		},
	}))
}
