package problem746

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMinCostClimbingStairs(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{10, 15, 20},
			want: 15,
		},
		{
			in:   []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}
	for _, tt := range tests {
		got := minCostClimbingStairs(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
