package two_sum

import (
	"testing"
)

type caseType struct {
	Nums     []int
	Target   int
	Expected []int
}

func TestTwoSum(t *testing.T) {
	tests := [...]caseType{
		{
			Nums:     []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{0, 1},
		},
		{
			Nums:     []int{2, 7, 11, 15},
			Target:   10,
			Expected: nil,
		},
		{
			Nums:     []int{1, 2, 3, 4},
			Target:   5,
			Expected: []int{1, 2},
		},
	}

	for i, test := range tests {
		output := twoSum(test.Nums, test.Target)
		if len(output) != len(test.Expected) {
			t.Fatalf(
				"Case: %v, Input: %v, Output: %v, Expected: %v",
				i, test.Nums, output, test.Expected,
			)
		}
		for k, v := range test.Expected {
			if output[k] != v {
				t.Fatalf(
					"Case: %v, Input: %v, Output: %v, Expected: %v",
					i, test.Nums, output, test.Expected,
				)
			}
		}
	}
}
