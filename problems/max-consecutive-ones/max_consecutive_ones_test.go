package problem485

import "testing"

type testType struct {
	in   []int
	want int
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 1, 0, 1, 1, 1},
			want: 3,
		},
		{
			in:   []int{1, 0, 1, 1, 0, 1},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := findMaxConsecutiveOnes(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
