package problem724

import "testing"

type testType struct {
	in   []int
	want int
}

func TestPivotIndex(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			in:   []int{1, 2, 3},
			want: -1,
		},
		{
			in:   []int{0},
			want: 0,
		},
		{
			in:   []int{0, 0},
			want: 0,
		},
		{
			in:   []int{1, 2, 1},
			want: 1,
		},
		{
			in:   []int{},
			want: -1,
		},
	}
	for _, tt := range tests {
		got := pivotIndex(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
