package problem1005

import "testing"

type testType struct {
	in   []int
	k    int
	want int
}

func TestLargestSumAfterKNegations(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{4, 2, 3},
			k:    1,
			want: 5,
		},
		{
			in:   []int{3, -1, 0, 2},
			k:    3,
			want: 6,
		},
		{
			in:   []int{2, -3, -1, 5, -4},
			k:    2,
			want: 13,
		},
	}
	for _, tt := range tests {
		got := largestSumAfterKNegations(tt.in, tt.k)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
