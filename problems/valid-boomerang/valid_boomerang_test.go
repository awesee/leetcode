package problem1037

import "testing"

type testType struct {
	in   [][]int
	want bool
}

func TestIsBoomerang(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 1},
				{2, 3},
				{3, 2},
			},
			want: true,
		},
		{
			in: [][]int{
				{1, 1},
				{2, 2},
				{3, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isBoomerang(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
