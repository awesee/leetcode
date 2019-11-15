package problem1013

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestCanThreePartsEqualSum(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			want: true,
		},
		{
			in:   []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			want: false,
		},
		{
			in:   []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
			want: true,
		},
		{
			in:   []int{-3, 3, 3, -3},
			want: false,
		},
		{
			in:   []int{1, 7, 2, 6, 3, 5, 8},
			want: false,
		},
	}
	for _, tt := range tests {
		got := canThreePartsEqualSum(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
