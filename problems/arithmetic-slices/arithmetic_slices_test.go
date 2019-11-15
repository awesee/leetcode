package problem413

import "testing"

type testType struct {
	in   []int
	want int
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 3, 5, 7, 9},
			want: 6,
		},
		{
			in:   []int{7, 7, 7, 7},
			want: 3,
		},
		{
			in:   []int{3, -1, -5, -9},
			want: 3,
		},
		{
			in:   []int{1, 2, 3, 4},
			want: 3,
		},
		{
			in:   []int{1, 2, 3, 4, 6, 8, 10},
			want: 6,
		},
		{
			in:   []int{1, 2},
			want: 0,
		},
		{
			in:   []int{1, 2, 4, 5, 7},
			want: 0,
		},
	}
	for _, tt := range tests {
		got := numberOfArithmeticSlices(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
