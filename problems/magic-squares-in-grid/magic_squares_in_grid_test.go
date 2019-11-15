package problem840

import "testing"

type testType struct {
	in   [][]int
	want int
}

func TestNumMagicSquaresInside(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{4, 3, 8, 4},
				{9, 5, 1, 9},
				{2, 7, 6, 2},
			},
			want: 1,
		},
		{
			in: [][]int{
				{5, 5, 5},
				{5, 5, 5},
				{5, 5, 5},
			},
			want: 0,
		},
		{
			in: [][]int{
				{1, 8, 6},
				{10, 5, 0},
				{4, 2, 9},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		got := numMagicSquaresInside(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
