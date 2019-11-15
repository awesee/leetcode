package problem581

import "testing"

type testType struct {
	in   []int
	want int
}

func TestFindUnsortedSubarray(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{2, 6, 4, 8, 10, 9, 15},
			want: 5,
		},
		{
			in:   []int{1, 2, 3, 4},
			want: 0,
		},
		{
			in:   []int{1, 2, 4, 5, 3},
			want: 3,
		},
	}
	for _, tt := range tests {
		got := findUnsortedSubarray(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
