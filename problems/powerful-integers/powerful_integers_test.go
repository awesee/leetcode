package powerful_integers

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	x        int
	y        int
	bound    int
	expected []int
}

func TestPowerfulIntegers(t *testing.T) {
	tests := [...]caseType{
		{
			x:        2,
			y:        3,
			bound:    10,
			expected: []int{2, 3, 4, 5, 7, 9, 10},
		},
		{
			x:        3,
			y:        5,
			bound:    15,
			expected: []int{2, 4, 6, 8, 10, 14},
		},
	}
	for _, tc := range tests {
		output := powerfulIntegers(tc.x, tc.y, tc.bound)
		if !IsEqualSliceInt(output, tc.expected) {
			t.Fatalf("input: %v %v %v, output: %v, expected: %v", tc.x, tc.y, tc.bound, output, tc.expected)
		}
	}
}
