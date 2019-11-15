package problem994

import "testing"

type testType struct {
	in   [][]int
	want int
}

func TestOrangesRotting(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			in: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			in: [][]int{
				{0, 2},
			},
			want: 0,
		},
		{
			in: [][]int{
				{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		got := orangesRotting(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
