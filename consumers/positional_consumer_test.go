package consumers

import (
	"testing"
)

type TestCase struct {
	args              []string
	expectedConsumed  []string
	expectedRemainder []string
}

var testCaseSuit = []TestCase{
	{
		args:              []string{"some-dir", "some-other-dir"},
		expectedConsumed:  []string{"some-dir", "some-other-dir"},
		expectedRemainder: []string{},
	},
}

func TestPositionalConsumer_Consume(t *testing.T) {
	for _, testCase := range testCaseSuit {
		consumer := PositionalConsumer{}
		args := testCase.args
		expectedConsumed := testCase.expectedConsumed
		expectedRemainder := testCase.expectedRemainder

		consumed, remainder := consumer.Consume(args)

		if !stringArraysEqual(expectedConsumed, consumed) {
			t.Errorf(
				"Expected consumed %#v, got %#v when consuming args %#v",
				expectedConsumed, consumed, args,
			)
		}

		if !stringArraysEqual(expectedRemainder, remainder) {
			t.Errorf(
				"Expected remainder %#v, got %#v when consuming args %#v",
				expectedRemainder, remainder, args,
			)
		}
	}
}

func stringArraysEqual(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, el := range a {
		if el != b[idx] {
			return false
		}
	}
	return true
}
