package utils

import (
	"testing"
)

type mapsEqualTestCase struct {
	mapA           map[string]string
	mapB           map[string]string
	expectedResult bool
}

var mapsEqualTestSuit = []mapsEqualTestCase{
	{
		map[string]string{"a": "1", "b": "2"},
		map[string]string{"a": "1", "b": "2"},
		true,
	},
	{
		map[string]string{"a": "1", "b": "2"},
		map[string]string{"a": "1", "b": "3"},
		false,
	},
	{
		map[string]string{"a": "1", "b": "2"},
		map[string]string{"a": "1", "c": "2"},
		false,
	},
	{
		map[string]string{"a": "1", "b": "2"},
		map[string]string{"a": "1"},
		false,
	},
	{
		map[string]string{},
		map[string]string{},
		true,
	},
	{
		map[string]string{"a": "1", "b": "2"},
		map[string]string{},
		false,
	},
}

func testMapsEqual(t *testing.T) {
	for _, testCase := range mapsEqualTestSuit {
		mapA := testCase.mapA
		mapB := testCase.mapB
		expectedResult := testCase.expectedResult

		result := StringMapsEqual(mapA, mapB)

		if result != expectedResult {
			t.Errorf(
				"Expected %#v, got %#v when comparing %#v for equality with %#v",
				expectedResult, result, mapA, mapB,
			)
		}
	}
}
