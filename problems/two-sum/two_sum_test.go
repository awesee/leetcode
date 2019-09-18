package problem_1

import (
	"reflect"
	"testing"
)

type caseType struct {
	nums     []int
	target   int
	expected []int
}

func TestTwoSum(t *testing.T) {
	tests := [...]caseType{
		{
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			nums:     []int{2, 7, 11, 15},
			target:   10,
			expected: nil,
		},
		{
			nums:     []int{1, 2, 3, 4},
			target:   5,
			expected: []int{1, 2},
		},
	}
	for _, tc := range tests {
		output := twoSum(tc.nums, tc.target)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.nums, output, tc.expected)
		}
	}
}
