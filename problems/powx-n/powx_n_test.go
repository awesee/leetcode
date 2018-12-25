package powx_n

import (
	"math"
	"testing"
)

type caseType struct {
	x        float64
	n        int
	expected float64
}

func TestMyPow(t *testing.T) {
	tests := [...]caseType{
		{
			x:        2.00000,
			n:        10,
			expected: 1024.00000,
		},
		{
			x:        2.10000,
			n:        3,
			expected: 9.26100,
		},
		{
			x:        2.00000,
			n:        -2,
			expected: 0.25000,
		},
	}

	for _, tc := range tests {
		output := myPow(tc.x, tc.n)
		output = math.Trunc(output*1e5) / 1e5
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.x, tc.n, output, tc.expected)
		}
	}
}
