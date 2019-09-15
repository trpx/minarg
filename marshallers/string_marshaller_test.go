package marshallers

import (
	"github.com/trpx/minarg/errors"
	"testing"
)

type stringMarshallerTestCase struct {
	value             string
	expectedResult    string
	expectedErrorCode int
}

var stringMarshallerTestCaseSuit = []stringMarshallerTestCase{
	{
		value:          "val",
		expectedResult: "val",
	},
	{
		value:          "",
		expectedResult: "",
	},
}

func TestStringMarshaller(t *testing.T) {
	for _, testCase := range stringMarshallerTestCaseSuit {
		marshaller := StringMarshaller{}
		value := testCase.value
		expectedResult := testCase.expectedResult
		expectedErrorCode := testCase.expectedErrorCode

		var result *string
		var err errors.ParseError

		result, err = marshaller.Marshall(value)

		if expectedErrorCode != 0 {
			if err == nil {
				t.Errorf(
					"Haven't got expected error when marshalling value %#v.",
					value,
				)
			} else if expectedErrorCode != err.Code() {
				t.Errorf(
					"Got unexpected error '%s' when marshalling value %#v, expected another error.",
					err, value,
				)
			}
			continue
		}

		if err != nil {
			t.Errorf(
				"Got unexpected error '%s' when marshalling value %#v, expected no error",
				err, value,
			)
			continue
		}

		if result != nil {
			if *result != expectedResult {
				t.Errorf(
					"Expected '%s', got '%s' as a result of marshalling value %#v",
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
