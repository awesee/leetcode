package problem674

import "testing"

type testType struct {
	in   []int
	want int
}

func TestFindLengthOfLCIS(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 3, 5, 4, 7},
			want: 3,
		},
		{
			in:   []int{2, 2, 2, 2, 2},
			want: 1,
		},
		{
			in:   []int{},
			want: 0,
		},
		{
			in:   []int{1, 3, 5, 7},
			want: 4,
		},
	}
	for _, tt := range tests {
		got := findLengthOfLCIS(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
