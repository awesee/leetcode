package problem532

import "testing"

type testType struct {
	nums []int
	k    int
	want int
}

func TestFindPairs(t *testing.T) {
	tests := [...]testType{
		{
			nums: []int{1, 3, 1, 5, 4},
			k:    0,
			want: 1,
		},
		{
			nums: []int{3, 1, 4, 1, 5},
			k:    2,
			want: 2,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			k:    1,
			want: 4,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			k:    -1,
			want: 0,
		},
	}
	for _, tt := range tests {
		got := findPairs(tt.nums, tt.k)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
