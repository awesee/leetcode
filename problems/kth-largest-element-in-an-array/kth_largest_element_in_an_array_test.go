package problem215

import "testing"

type testType struct {
	in   []int
	k    int
	want int
}

func TestFindKthLargest(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		}, {
			in:   []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
	}
	for _, tt := range tests {
		got := findKthLargest(tt.in, tt.k)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
