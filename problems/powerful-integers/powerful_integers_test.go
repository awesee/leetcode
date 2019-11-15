package problem970

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	x     int
	y     int
	bound int
	want  []int
}

func TestPowerfulIntegers(t *testing.T) {
	tests := [...]testType{
		{
			x:     2,
			y:     3,
			bound: 10,
			want:  []int{2, 3, 4, 5, 7, 9, 10},
		},
		{
			x:     3,
			y:     5,
			bound: 15,
			want:  []int{2, 4, 6, 8, 10, 14},
		},
	}
	for _, tt := range tests {
		got := powerfulIntegers(tt.x, tt.y, tt.bound)
		if !kit.IsEqualSliceInt(got, tt.want) {
			t.Fatalf("in: %v %v %v, got: %v, want: %v", tt.x, tt.y, tt.bound, got, tt.want)
		}
	}
}
