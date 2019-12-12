package problem453

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMinMoves(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3},
			want: 3,
		},
		{
			in:   []int{5, 3, 1},
			want: 6,
		},
	}
	for _, tt := range tests {
		got := minMoves(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
