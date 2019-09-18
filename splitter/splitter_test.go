package splitter

import (
	"github.com/trpx/minarg/utils"
	"testing"
)

type splitterTestCase struct {
	args              []string
	splitter          Splitter
	expectedParts     Parts
	expectedRemainder []string
}

var splitterTestSuit = []splitterTestCase{
	{
		args: []string{"pos1", "pos2", "-abc", "--ab", "ab-val", "--cd", "cd-val", "comm1", "com_arg1", "--com-kw1"},
		splitter: Splitter{
			stringArgNames:  []string{"ab", "cd"},
			boolArgNames:    []string{"a", "b", "c"},
			subCommandNames: []string{"comm1"},
		},
		expectedParts: Parts{
			Positional: []string{"pos1", "pos2"},
			NamedStrings: map[string][]string{
				"cd": []string{"cd-val"},
				"ab": []string{"ab-val"},
			},
			NamedBool: map[string][]bool{
				"a": []bool{true},
				"b": []bool{true},
				"c": []bool{true},
			},
			Subcommand: []string{"comm1", "com_arg1", "--com-kw1"},
		},
		expectedRemainder: []string{},
	},
	{
		args: []string{"pos1", "pos2", "-abc", "--ab", "--cd", "cd-val", "comm1", "com_arg1", "--com-kw1"},
		splitter: Splitter{
			stringArgNames:  []string{"cd"},
			boolArgNames:    []string{"a", "b", "c", "ab"},
			subCommandNames: []string{"comm1"},
		},
		expectedParts: Parts{
			Positional: []string{"pos1", "pos2"},
			NamedStrings: map[string][]string{
				"cd": []string{"cd-val"},
			},
			NamedBool: map[string][]bool{
				"a":  []bool{true},
				"b":  []bool{true},
				"c":  []bool{true},
				"ab": []bool{true},
			},
			Subcommand: []string{"comm1", "com_arg1", "--com-kw1"},
		},
		expectedRemainder: []string{},
	},
	{
		args: []string{"pos1", "pos2", "-abc", "--ab", "--cd", "cd-val", "comm1", "com_arg1", "--com-kw1"},
		splitter: Splitter{
			stringArgNames:  []string{"ab", "cd"},
			boolArgNames:    []string{"a", "b", "c"},
			subCommandNames: []string{"comm1"},
		},
		expectedParts: Parts{
			Positional: []string{"pos1", "pos2"},
			NamedStrings: map[string][]string{
				"ab": []string{"--cd"},
			},
			NamedBool: map[string][]bool{
				"a": []bool{true},
				"b": []bool{true},
				"c": []bool{true},
			},
			Subcommand: []string{},
		},
		expectedRemainder: []string{"cd-val", "comm1", "com_arg1", "--com-kw1"},
	},
}

func TestPositionalConsumer_Consume(t *testing.T) {
	for _, testCase := range splitterTestSuit {

		splitter := testCase.splitter
		args := testCase.args

		expectedParts := testCase.expectedParts
		expectedRemainder := testCase.expectedRemainder

		expectedNamedStrings := expectedParts.NamedStrings
		expectedNamedBool := expectedParts.NamedBool
		expectedPositional := expectedParts.Positional
		expectedSubcommand := expectedParts.Subcommand

		parts, remainder := splitter.Split(args)

		namedBool := parts.NamedBool
		namedStrings := parts.NamedStrings
		positional := parts.Positional
		subcommand := parts.Subcommand

		if !utils.StringArraysEqual(expectedPositional, positional) {
			t.Errorf(
				"Expected positional %#v, got %#v when splitting args %#v",
				expectedPositional, positional, args,
			)
		}

		if !namedStringsEqual(expectedNamedStrings, namedStrings) {
			t.Errorf(
				"Expected named strings %#v, got %#v when splitting args %#v",
				expectedNamedStrings, namedStrings, args,
			)
		}

		if !namedBoolEqual(expectedNamedBool, namedBool) {
			t.Errorf(
				"Expected named bool %#v, got %#v when splitting args %#v",
				expectedNamedBool, namedBool, args,
			)
		}

		if !utils.StringArraysEqual(expectedSubcommand, subcommand) {
			t.Errorf(
				"Expected subcommand %#v, got %#v when splitting args %#v",
				expectedSubcommand, subcommand, args,
			)
		}

		if !utils.StringArraysEqual(expectedRemainder, remainder) {
			t.Errorf(
				"Expected remainder %#v, got %#v when splitting args %#v",
				expectedRemainder, remainder, args,
			)
		}
	}
}

func namedStringsEqual(argA map[string][]string, argB map[string][]string) bool {
	if len(argA) != len(argB) {
		return false
	}
	for k, v := range argA {
		bV, ok := argB[k]
		if !ok {
			return false
		}
		if !utils.StringArraysEqual(v, bV) {
			return false
		}
	}
	return true
}

func namedBoolEqual(argA map[string][]bool, argB map[string][]bool) bool {
	if len(argA) != len(argB) {
		return false
	}
	for k, v := range argA {
		bV, ok := argB[k]
		if !ok {
			return false
		}
		if len(v) != len(bV) {
			return false
		}
		for idx, i := range v {
			if i != bV[idx] {
				return false
			}
		}
	}
	return true
}
