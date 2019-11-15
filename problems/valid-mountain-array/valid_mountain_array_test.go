package problem941

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestValidMountainArray(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{2, 1},
			want: false,
		},
		{
			in:   []int{3, 5, 5},
			want: false,
		},
		{
			in:   []int{0, 3, 2, 1},
			want: true,
		},
		{
			in:   []int{1, 2, 3, 2},
			want: true,
		},
		{
			in:   []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			want: false,
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			want: false,
		},
	}
	for _, tt := range tests {
		got := validMountainArray(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
