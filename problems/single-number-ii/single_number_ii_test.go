package problem137

import "testing"

type testType struct {
	in   []int
	want int
}

func TestSingleNumber(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{2, 2, 3, 2},
			want: 3,
		},
		{
			in:   []int{0, 1, 0, 1, 0, 1, 99},
			want: 99,
		},
	}
	for _, tt := range tests {
		got := singleNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
