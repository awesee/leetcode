package problem697

import "testing"

type testType struct {
	in   []int
	want int
}

func TestFindShortestSubArray(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 2, 3, 1},
			want: 2,
		},
		{
			in:   []int{1, 2, 2, 3, 1, 4, 2},
			want: 6,
		},
		{
			in:   []int{1},
			want: 1,
		},
		{
			in:   []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2},
			want: 7,
		},
		{
			in:   []int{1, 2, 2, 3, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := findShortestSubArray(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
