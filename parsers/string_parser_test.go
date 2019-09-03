package parsers

import (
	"testing"
)

type TestCase struct {
	args           []string
	expectedResult string
	erroneous      bool
}

var testCaseSuit = []TestCase{
	{
		args:           []string{"--watch", "some-dir"},
		expectedResult: "some-dir",
	},
}

func TestStringParser(t *testing.T) {
	for _, testCase := range testCaseSuit {
		parser := StringParser{}
		args := testCase.args
		expected := testCase.expectedResult

		value, err := parser.parse(args)

		if value != expected {
			t.Errorf(
				"Expected '%s', got '%s' as a result of parsing args %#v",
				expected, value, args,
			)
		}

		if (err != nil) != testCase.erroneous {
			if err != nil {
				t.Errorf(
					"Got unexpected error %#v when parsing args %#v",
					err, args,
				)
			} else {
				t.Errorf(
					"Haven't got expected error when parsing args %#v",
					args,
				)
			}
			continue
		}
	}
}
