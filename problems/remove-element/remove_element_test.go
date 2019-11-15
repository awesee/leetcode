package problem27

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	val  int
	want []int
}

func TestRemoveElement(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 2, 2, 3},
			val:  3,
			want: []int{2, 2},
		},
		{
			in:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: []int{0, 1, 3, 0, 4},
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			val:  6,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			in:   []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4},
			val:  4,
			want: []int{1, 2, 2, 3, 3, 3},
		},
		{
			in:   []int{},
			val:  1,
			want: []int{},
		},
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.in))
		copy(nums, tt.in)
		l := removeElement(nums, tt.val)
		got := nums[:l]
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
