package problem561

import "testing"

type testType struct {
	in   []int
	want int
}

func TestArrayPairSum(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 4, 3, 2},
			want: 4,
		},
	}
	for _, tt := range tests {
		got := arrayPairSum(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
