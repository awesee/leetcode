package problem53

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMaxSubArray(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
	}
	for _, tt := range tests {
		got := maxSubArray(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
