package problem456

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestFind132pattern(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 4},
			want: false,
		},
		{
			in:   []int{3, 1, 4, 2},
			want: true,
		},
		{
			in:   []int{-1, 3, 2, 0},
			want: true,
		},
	}
	for _, tt := range tests {
		got := find132pattern(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
