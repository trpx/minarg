package accumulators

import (
	"github.com/trpx/minarg/errors"
	"github.com/trpx/minarg/utils"
	"testing"
)

type stringAccumulatorTestCase struct {
	capacity          int
	values            []string
	expectedResult    []string
	expectedErrorCode int
}

var stringAccumulatorTestCaseSuit = []stringAccumulatorTestCase{
	{
		capacity: 1,
		values:   []string{},
	},
	{
		capacity:       2,
		values:         []string{"val1", "val2"},
		expectedResult: []string{"val1", "val2"},
	},
	{
		capacity:          1,
		values:            []string{"val1", "val2"},
		expectedErrorCode: errors.TOO_MANY_VALUES,
	},
}

func TestStringAccumulator(t *testing.T) {
	for _, testCase := range stringAccumulatorTestCaseSuit {
		capacity := testCase.capacity
		accumulator := FixedAccumulator{Capacity: capacity}
		values := testCase.values
		expectedResult := testCase.expectedResult
		expectedErrorCode := testCase.expectedErrorCode

		var err errors.ParseError

		var result []string

		for _, val := range values {
			err = accumulator.Accumulate(val)
			result = accumulator.Result()
		}

		if expectedErrorCode != 0 {
			if err == nil {
				t.Errorf(
					"Haven't got expected error when accumulating values %#v with capacity %d.",
					values, capacity,
				)
			} else if expectedErrorCode != err.Code() {
				t.Errorf(
					"Got unexpected error '%s' when accumulating values %#v, expected another error.",
					err, values,
				)
			}
			continue
		}

		if err != nil {
			t.Errorf(
				"Got unexpected error '%s' when accumulating values %#v, expected no error",
				err, values,
			)
			continue
		}

		if !utils.StringArraysEqual(result, expectedResult) {
			t.Errorf(
				"Expected %#v, got %#v as a result of accumulating values %#v",
				expectedResult, result, values,
			)
		}
	}
}
