package api

import "testing"

func TestAnalysisParam_Validate(t *testing.T) {
	tests := []struct {
		param       AnalysisParam
		expectError bool
		description string
	}{
		{
			param:       AnalysisParam{Duration: "1h", Dimension: Comments},
			expectError: false,
			description: "Valid duration and dimension",
		},
		{
			param:       AnalysisParam{Duration: "invalid_duration", Dimension: Likes},
			expectError: true,
			description: "Invalid duration",
		},
		{
			param:       AnalysisParam{Duration: "1h", Dimension: "InvalidDimension"},
			expectError: true,
			description: "Invalid dimension",
		},
	}

	for _, test := range tests {
		err := test.param.Validate()

		if test.expectError && err == nil {
			t.Errorf("%s: Expected an error, but got none", test.description)
		}

		if !test.expectError && err != nil {
			t.Errorf("%s: Unexpected error: %v", test.description, err)
		}
	}
}
