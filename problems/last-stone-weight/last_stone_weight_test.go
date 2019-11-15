package problem1046

import "testing"

type testType struct {
	in   []int
	want int
}

func TestLastStoneWeight(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{2, 7, 4, 1, 8, 1},
			want: 1,
		},
		{
			in:   []int{2, 7, 4, 1, 8, 1, 5},
			want: 0,
		},
		{
			in:   []int{316, 157, 73, 106, 771, 828},
			want: 37,
		},
	}
	for _, tt := range tests {
		got := lastStoneWeight(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
