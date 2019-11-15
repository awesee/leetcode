package problem908

import "testing"

type testType struct {
	in   []int
	k    int
	want int
}

func TestSmallestRangeI(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1},
			k:    0,
			want: 0,
		},
		{
			in:   []int{0, 10},
			k:    2,
			want: 6,
		},
		{
			in:   []int{1, 3, 6},
			k:    3,
			want: 0,
		},
		{
			in:   []int{1, 3, 5, 7, 9},
			k:    1,
			want: 6,
		},
		{
			in:   []int{18, 16, 12, 7, 9, 3, 5},
			k:    6,
			want: 3,
		},
	}
	for _, tt := range tests {
		got := smallestRangeI(tt.in, tt.k)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
