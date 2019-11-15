package problem26

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestRemoveDuplicates(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			in:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			in:   []int{1, 2, 3, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			in:   []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			in:   []int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.in))
		copy(nums, tt.in)
		l := removeDuplicates(nums)
		got := nums[:l]
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
