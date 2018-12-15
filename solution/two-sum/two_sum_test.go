package two_sum

import (
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

	for _, test := range tests {
		output := twoSum(test.nums, test.target)
		if len(output) != len(test.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", test.nums, output, test.expected)
		}
		for k, v := range test.expected {
			if output[k] != v {
				t.Fatalf("input: %v, output: %v, expected: %v", test.nums, output, test.expected)
			}
		}
	}
}
