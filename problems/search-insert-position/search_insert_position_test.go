package problem35

import "testing"

type testType struct {
	nums   []int
	target int
	want   int
}

func TestSearchInsert(t *testing.T) {
	tests := [...]testType{
		{
			nums:   []int{2, 7, 11, 15},
			target: 5,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		got := searchInsert(tt.nums, tt.target)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.nums, got, tt.want)
		}
	}
}
