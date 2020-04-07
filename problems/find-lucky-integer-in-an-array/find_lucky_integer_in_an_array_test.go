package problem1394

import "testing"

type testType struct {
	in   []int
	want int
}

func TestFindLucky(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{2, 2, 3, 4},
			want: 2,
		},
		{
			in:   []int{1, 2, 2, 3, 3, 3},
			want: 3,
		},
		{
			in:   []int{2, 2, 2, 3, 3},
			want: -1,
		},
		{
			in:   []int{5},
			want: -1,
		},
		{
			in:   []int{7, 7, 7, 7, 7, 7, 7},
			want: 7,
		},
	}
	for _, tt := range tests {
		got := findLucky(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
