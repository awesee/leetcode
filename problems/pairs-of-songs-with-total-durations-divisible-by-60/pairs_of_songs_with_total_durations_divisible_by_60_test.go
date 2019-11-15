package problem1010

import "testing"

type testType struct {
	in   []int
	want int
}

func TestNumPairsDivisibleBy60(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{30, 20, 150, 100, 40},
			want: 3,
		},
		{
			in:   []int{60, 60, 60},
			want: 3,
		},
	}
	for _, tt := range tests {
		got := numPairsDivisibleBy60(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
