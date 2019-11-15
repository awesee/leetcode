package problem50

import (
	"math"
	"testing"
)

type testType struct {
	x    float64
	n    int
	want float64
}

func TestMyPow(t *testing.T) {
	tests := [...]testType{
		{
			x:    2.00000,
			n:    10,
			want: 1024.00000,
		},
		{
			x:    2.10000,
			n:    3,
			want: 9.26100,
		},
		{
			x:    2.00000,
			n:    -2,
			want: 0.25000,
		},
	}

	for _, tt := range tests {
		got := myPow(tt.x, tt.n)
		got = math.Trunc(got*1e5) / 1e5
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.x, tt.n, got, tt.want)
		}
	}
}
