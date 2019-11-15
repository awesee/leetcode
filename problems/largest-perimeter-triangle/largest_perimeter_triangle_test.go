package problem976

import "testing"

type testType struct {
	in   []int
	want int
}

func TestLargestPerimeter(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{2, 1, 2},
			want: 5,
		},
		{
			in:   []int{1, 2, 1},
			want: 0,
		},
		{
			in:   []int{3, 2, 3, 4},
			want: 10,
		},
		{
			in:   []int{3, 6, 2, 3},
			want: 8,
		},
	}
	for _, tt := range tests {
		got := largestPerimeter(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
