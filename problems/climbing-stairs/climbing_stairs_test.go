package problem70

import "testing"

type testType struct {
	in   int
	want int
}

func TestClimbStairs(t *testing.T) {
	tests := [...]testType{
		{
			in:   1,
			want: 1,
		},
		{
			in:   2,
			want: 2,
		},
		{
			in:   3,
			want: 3,
		},
	}
	for _, tt := range tests {
		got := climbStairs(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
