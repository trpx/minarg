package parsers

import (
	"github.com/trpx/minarg/errors"
	"testing"
)

type stringParserTestCase struct {
	value             string
	expectedResult    string
	expectedErrorCode int
}

var stringParserTestCaseSuit = []stringParserTestCase{
	{
		value:          "val",
		expectedResult: "val",
	},
	{
		value:          "",
		expectedResult: "",
	},
}

func TestStringParser(t *testing.T) {
	for _, testCase := range stringParserTestCaseSuit {
		parser := StringParser{}
		value := testCase.value
		expectedResult := testCase.expectedResult
		expectedErrorCode := testCase.expectedErrorCode

		var result *string
		var err errors.ParseError

		result, err = parser.parse(value)

		if expectedErrorCode != 0 {
			if err == nil {
				t.Errorf(
					"Haven't got expected error when parsing value %#v.",
					value,
				)
			} else if expectedErrorCode != err.Code() {
				t.Errorf(
					"Got unexpected error '%s' when parsing value %#v, expected another error.",
					err, value,
				)
			}
			continue
		}

		if err != nil {
			t.Errorf(
				"Got unexpected error '%s' when parsing value %#v, expected no error",
				err, value,
			)
			continue
		}

		if result != nil {
			if *result != expectedResult {
				t.Errorf(
					"Expected '%s', got '%s' as a result of parsing value %#v",
					expectedResult, *result, value,
				)
			}
		} else {
			t.Errorf(
				"Test implementation error, result is nil, but shouldn't"+
					" ever be here, value: %#v, expected result: '%s'.",
				value, expectedResult,
			)
		}
	}
}
