package problem219

import "testing"

type testType struct {
	nums []int
	k    int
	want bool
}

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := [...]testType{
		{
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
	}
	for _, tt := range tests {
		got := containsNearbyDuplicate(tt.nums, tt.k)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
