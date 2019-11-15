package problem1051

import "testing"

type testType struct {
	in   []int
	want int
}

func TestHeightChecker(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 1, 4, 2, 1, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		got := heightChecker(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
