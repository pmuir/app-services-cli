package validation

import (
	"testing"

	"github.com/redhat-developer/app-services-cli/internal/localizer"
)

func TestValidateName(t *testing.T) {
	_ = localizer.IncludeAssetsAndLoadMessageFiles()

	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "fails when length is 0",
			args:    args{""},
			wantErr: true,
		},
		{
			name:    "passes when length is 1",
			args:    args{"s"},
			wantErr: false,
		},
		{
			name:    "passes when length is 50 (max length)",
			args:    args{"ssssssssssssssssssssssssssssssssssssssssssssssssss"},
			wantErr: false,
		},
		{
			name:    "fails when length exceeds max length",
			args:    args{"sssssssssssssssssssssssssssssssssssssssssssssssssss"},
			wantErr: true,
		},
		{
			name:    "passes on valid name",
			args:    args{"svcacctone"},
			wantErr: false,
		},
		{
			name:    "passes on valid name with hyphens",
			args:    args{"svc-acct-one"},
			wantErr: false,
		},
		{
			name:    "passes on valid name with digits",
			args:    args{"svc-acct-1s"},
			wantErr: false,
		},
		{
			name:    "fails with capital letters",
			args:    args{"Svc-acct-one"},
			wantErr: true,
		},
		{
			name:    "fails number in first section",
			args:    args{"1svc-acct-one"},
			wantErr: true,
		},
	}
	// nolint:scopelint
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateName(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDescription(t *testing.T) {
	_ = localizer.IncludeAssetsAndLoadMessageFiles()

	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "passes when empty",
			args:    args{""},
			wantErr: false,
		},
		{
			name:    "passes on max length (255)",
			args:    args{"trl1rmcyl6dp4xxqy0rwudhodbpjc4crja8ibf2yco6obalko6qor9n2a1wsqruolg0ewrndumw2xkezzuwg8pjo6ntsmi1cjw99hjcko4t2kjkxmaswzgk8ko75pcs4js0pzypuyjxxnld4dijxadzs8peioi6d5jjxxtfl9vicufmxuacvu7m8ycbwhsbiu9ipw5fxplf0ojs8bxd7hwt4rn4phbcdgivxdzprhyfjamkgjzytjz25cmqagtw"},
			wantErr: false,
		},
		{
			name:    "fails when exceeds max length",
			args:    args{"trl1rmcyl6dp4xxqy0rwudhodbpjc4crja8ibf2yco6obalko6qor9n2a1wsqruolg0ewrndumw2xkezzuwg8pjo6ntsmi1cjw99hjcko4t2kjkxmaswzgk8ko75pcs4js0pzypuyjxxnld4dijxadzs8peioi6d5jjxxtfl9vicufmxuacvu7m8ycbwhsbiu9ipw5fxplf0ojs8bxd7hwt4rn4phbcdgivxdzprhyfjamkgjzytjz25cmqagtwa"},
			wantErr: true,
		},
		{
			name:    "passes with spaces",
			args:    args{"here is a description"},
			wantErr: false,
		},
		{
			name:    "fails with special character",
			args:    args{"here is a description!"},
			wantErr: true,
		},
		{
			name:    "passes with capital letters",
			args:    args{"Hello"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		// nolint:scopelint
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateDescription(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("ValidateDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
