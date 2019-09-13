package consumers

import (
	"github.com/trpx/minarg/utils"
	"testing"
)

type stringNamedConsumerTestCase struct {
	consumerArgName   string
	args              []string
	expectedConsumed  []string
	expectedRemainder []string
}

var stringNamedConsumerTestCaseSuit = []stringNamedConsumerTestCase{
	{
		consumerArgName:   "kw1",
		args:              []string{"--kw1", "val1", "--kw2", "val2", "pos1"},
		expectedConsumed:  []string{"val1"},
		expectedRemainder: []string{"--kw2", "val2", "pos1"},
	},
	{
		consumerArgName:   "pos1",
		args:              []string{"pos1", "pos2"},
		expectedConsumed:  []string{},
		expectedRemainder: []string{"pos1", "pos2"},
	},
	{
		consumerArgName:   "kw1",
		args:              []string{"pos1", "pos2", "--kw1", "val1"},
		expectedConsumed:  []string{},
		expectedRemainder: []string{"pos1", "pos2", "--kw1", "val1"},
	},
	{
		consumerArgName:   "kw1",
		args:              []string{"--kw1", "--kw2"},
		expectedConsumed:  []string{"--kw2"},
		expectedRemainder: []string{},
	},
	{
		consumerArgName:   "kw2",
		args:              []string{"--kw1", "--kw2"},
		expectedConsumed:  []string{},
		expectedRemainder: []string{"--kw1", "--kw2"},
	},
	{
		consumerArgName:   "pos1",
		args:              []string{},
		expectedConsumed:  []string{},
		expectedRemainder: []string{},
	},
}

func TestStringNamedConsumer_Consume(t *testing.T) {
	for _, testCase := range stringNamedConsumerTestCaseSuit {
		consumer := StringNamedConsumer{testCase.consumerArgName}
		args := testCase.args
		expectedConsumed := testCase.expectedConsumed
		expectedRemainder := testCase.expectedRemainder

		consumed, remainder := consumer.Consume(args)

		if !utils.StringArraysEqual(expectedConsumed, consumed) {
			t.Errorf(
				"Expected consumed %#v, got %#v when consuming args %#v",
				expectedConsumed, consumed, args,
			)
		}

		if !utils.StringArraysEqual(expectedRemainder, remainder) {
			t.Errorf(
				"Expected remainder %#v, got %#v when consuming args %#v",
				expectedRemainder, remainder, args,
			)
		}
	}
}
