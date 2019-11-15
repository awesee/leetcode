package problem643

import "testing"

type testType struct {
	nums []int
	k    int
	want float64
}

func TestFindMaxAverage(t *testing.T) {
	tests := [...]testType{
		{
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			want: 12.75,
		},
	}
	for _, tt := range tests {
		got := findMaxAverage(tt.nums, tt.k)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.nums, tt.k, got, tt.want)
		}
	}
}
