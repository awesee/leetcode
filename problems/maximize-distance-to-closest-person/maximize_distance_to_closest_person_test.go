package problem849

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMaxDistToClosest(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 0, 0, 0, 1, 0, 1},
			want: 2,
		},
		{
			in:   []int{1, 0, 0, 0},
			want: 3,
		},
		{
			in:   []int{0, 0, 0, 1},
			want: 3,
		},
		{
			in:   []int{0, 0, 0, 0, 0, 1, 0},
			want: 5,
		},
		{
			in:   []int{0, 1, 0, 0, 0, 1, 0},
			want: 2,
		},
		{
			in:   []int{0, 1, 0, 0, 0, 0, 0},
			want: 5,
		},
		{
			in:   []int{1, 0, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		got := maxDistToClosest(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
