package problem628

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMaximumProduct(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3},
			want: 6,
		},
		{
			in:   []int{1, 2, 3, 4},
			want: 24,
		},
		{
			in:   []int{-2, -1, 0, 2, 3},
			want: 6,
		},
		{
			in:   []int{-1, -2, -3, 5, 4, 3, 2, 1},
			want: 60,
		},
	}
	for _, tt := range tests {
		got := maximumProduct(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
