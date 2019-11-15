package problem747

import "testing"

type testType struct {
	in   []int
	want int
}

func TestDominantIndex(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 6, 1, 0},
			want: 1,
		},
		{
			in:   []int{1, 2, 3, 4},
			want: -1,
		},
		{
			in:   []int{1},
			want: 0,
		},
		{
			in:   []int{0, 0, 3, 2},
			want: -1,
		},
		{
			in:   []int{3, 0, 0, 2},
			want: -1,
		},
	}
	for _, tt := range tests {
		got := dominantIndex(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
