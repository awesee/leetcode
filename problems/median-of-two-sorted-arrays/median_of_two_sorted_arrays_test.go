package median_of_two_sorted_arrays

import "testing"

type caseType struct {
	num1     []int
	num2     []int
	expected float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	tests := [...]caseType{
		{
			num1:     []int{1, 3},
			num2:     []int{2},
			expected: 2.0,
		},
		{
			num1:     []int{1, 2},
			num2:     []int{3, 4},
			expected: 2.5,
		},
		{
			num1:     []int{1, 3, 5, 7},
			num2:     []int{2, 4},
			expected: 3.5,
		},
		{
			num1:     []int{1, 3, 5},
			num2:     []int{},
			expected: 3,
		},
		{
			num1:     []int{1, 2},
			num2:     []int{3},
			expected: 2,
		},
		{
			num1:     []int{1, 1},
			num2:     []int{1, 1},
			expected: 1,
		},
		{
			num1:     []int{1, 1, 1},
			num2:     []int{1, 1},
			expected: 1,
		},
		{
			num1:     []int{1, 1},
			num2:     []int{1, 1, 1},
			expected: 1,
		},
		{
			num1:     []int{1},
			num2:     []int{2, 3, 4},
			expected: 2.5,
		},
		{
			num1:     []int{2, 3, 4},
			num2:     []int{1},
			expected: 2.5,
		},
		{
			num1:     []int{},
			num2:     []int{1},
			expected: 1,
		},
		{
			num1:     []int{},
			num2:     []int{1, 1},
			expected: 1,
		},
	}

	for _, tc := range tests {
		output := findMedianSortedArrays(tc.num1, tc.num2)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.num1, tc.num2, output, tc.expected)
		}
	}
}
