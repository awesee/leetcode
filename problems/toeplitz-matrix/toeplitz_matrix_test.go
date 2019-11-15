package problem766

import "testing"

type testType struct {
	in   [][]int
	want bool
}

func TestIsToeplitzMatrix(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			want: true,
		},
		{
			in: [][]int{
				{1, 2},
				{2, 2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isToeplitzMatrix(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
