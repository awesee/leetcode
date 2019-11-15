package problem11

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMaxArea(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want: 49,
		},
		{
			in:   []int{1, 8, 6, 30, 20, 6, 9, 10, 1},
			want: 48,
		},
	}
	for _, tt := range tests {
		got := maxArea(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
