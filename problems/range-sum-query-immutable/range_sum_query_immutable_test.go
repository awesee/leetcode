package range_sum_query_immutable

import "testing"

type caseType struct {
	i        int
	j        int
	expected int
}

func TestConstructor(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	tests := [...]caseType{
		{
			i:        0,
			j:        2,
			expected: 1,
		},
		{
			i:        2,
			j:        5,
			expected: -1,
		},
		{
			i:        0,
			j:        5,
			expected: -3,
		},
	}
	for _, tc := range tests {
		output := obj.SumRange(tc.i, tc.j)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.i, tc.j, output, tc.expected)
		}
	}
}
