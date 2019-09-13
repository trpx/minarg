package utils

import (
	"testing"
)

type TestCase struct {
	arrayA         []string
	arrayB         []string
	expectedResult bool
}

var testCaseSuit = []TestCase{
	{
		[]string{"a", "b"},
		[]string{"a", "b"},
		true,
	},
	{
		[]string{"a", "b"},
		[]string{"a", "c"},
		false,
	},
	{
		[]string{"a", "b"},
		[]string{"a"},
		false,
	},
	{
		[]string{},
		[]string{},
		true,
	},
	{
		[]string{"a", "b"},
		[]string{},
		false,
	},
}

func TestPositionalConsumer_Consume(t *testing.T) {
	for _, testCase := range testCaseSuit {
		arrayA := testCase.arrayA
		arrayB := testCase.arrayB
		expectedResult := testCase.expectedResult

		result := StringArraysEqual(arrayA, arrayB)

		if result != expectedResult {
			t.Errorf(
				"Expected %#v, got %#v when comparing %#v for equality with %#v",
				expectedResult, result, arrayA, arrayB,
			)
		}
	}
}
