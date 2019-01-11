package fair_candy_swap

import (
	"reflect"
	"testing"
)

type caseType struct {
	a        []int
	b        []int
	expected []int
}

func TestFairCandySwap(t *testing.T) {
	tests := [...]caseType{
		{
			a:        []int{1, 1},
			b:        []int{2, 2},
			expected: []int{1, 2},
		},
		{
			a:        []int{1, 2},
			b:        []int{2, 3},
			expected: []int{1, 2},
		},
		{
			a:        []int{2},
			b:        []int{1, 3},
			expected: []int{2, 3},
		},
		{
			a:        []int{1, 2, 5},
			b:        []int{2, 4},
			expected: []int{5, 4},
		},
		{
			a:        []int{1, 2, 3},
			b:        []int{11, 15, 17},
			expected: nil,
		},
	}
	for _, tc := range tests {
		output := fairCandySwap(tc.a, tc.b)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.a, tc.b, output, tc.expected)
		}
	}
}
