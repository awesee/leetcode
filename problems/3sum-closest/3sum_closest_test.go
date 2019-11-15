package problem16

import "testing"

type testType struct {
	in     []int
	target int
	want   int
}

func TestThreeSumClosest(t *testing.T) {
	tests := [...]testType{
		{
			in:     []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			in:     []int{-1, 0, 1, 2, -1, -4},
			target: 1,
			want:   1,
		},
		{
			in:     []int{0, 0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			in:     []int{-2, 0, 0, 2, 2, 2},
			target: 2,
			want:   2,
		},
		{
			in:     []int{-2, 0, 0, 2, 2, 2, 2},
			target: 1,
			want:   0,
		},
	}
	for _, tt := range tests {
		got := threeSumClosest(tt.in, tt.target)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
