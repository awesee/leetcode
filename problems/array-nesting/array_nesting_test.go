package problem565

import (
	"testing"
)

type testType struct {
	in   []int
	want int
}

func TestArrayNesting(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{5, 4, 0, 3, 1, 6, 2},
			want: 4,
		},
		{
			in:   []int{0, 3, 1, 5, 4, 6, 2},
			want: 5,
		},
		{
			in:   []int{6, 2, 5, 4, 0, 3, 1},
			want: 7,
		},
		{
			in:   []int{3, 1, 5, 4, 0, 6, 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		got := arrayNesting(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
