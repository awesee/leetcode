package problem914

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestHasGroupsSizeX(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 4, 4, 3, 2, 1},
			want: true,
		},
		{
			in:   []int{1, 1, 1, 2, 2, 2, 3, 3},
			want: false,
		},
		{
			in:   []int{1},
			want: false,
		},
		{
			in:   []int{1, 1},
			want: true,
		},
		{
			in:   []int{1, 1, 2, 2, 2, 2},
			want: true,
		},
	}
	for _, tt := range tests {
		got := hasGroupsSizeX(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
