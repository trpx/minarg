package consumers

import (
	"github.com/trpx/minarg/utils"
	"testing"
)

type TestCase struct {
	args              []string
	expectedConsumed  []string
	expectedRemainder []string
}

var testCaseSuit = []TestCase{
	{
		args:              []string{"arg1", "arg2"},
		expectedConsumed:  []string{"arg1", "arg2"},
		expectedRemainder: []string{},
	},
	{
		args:              []string{"arg1", "arg2", "--named-arg1", "--named-arg2"},
		expectedConsumed:  []string{"arg1", "arg2"},
		expectedRemainder: []string{"--named-arg1", "--named-arg2"},
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

func TestPositionalConsumer_Consume(t *testing.T) {
	for _, testCase := range testCaseSuit {
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
