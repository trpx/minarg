package consumers

import (
	"github.com/trpx/minarg/utils"
	"testing"
)

type stringNamedConsumerTestCase struct {
	args              []string
	expectedConsumed  []string
	expectedRemainder []string
}

var stringNamedConsumerTestCaseSuit = []stringNamedConsumerTestCase{
	{
		args:              []string{"--named-arg1", "value1", "--named-arg2", "value2", "arg1"},
		expectedConsumed:  []string{"--named-arg1", "value1"},
		expectedRemainder: []string{"--named-arg2", "value2", "arg1"},
	},
	{
		args:              []string{"arg1", "arg2"},
		expectedConsumed:  []string{},
		expectedRemainder: []string{"arg1", "arg2"},
	},
	{
		args:              []string{"arg1", "arg2", "--named-arg1", "--named-arg2"},
		expectedConsumed:  []string{},
		expectedRemainder: []string{"arg1", "arg2", "--named-arg1", "--named-arg2"},
	},
	{
		args:              []string{"--named-arg1", "--named-arg2"},
		expectedConsumed:  []string{},
		expectedRemainder: []string{"--named-arg1", "--named-arg2"},
	},
	{
		args:              []string{},
		expectedConsumed:  []string{},
		expectedRemainder: []string{},
	},
}

func TestStringNamedConsumer_Consume(t *testing.T) {
	for _, testCase := range stringNamedConsumerTestCaseSuit {
		consumer := PositionalConsumer{}
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
