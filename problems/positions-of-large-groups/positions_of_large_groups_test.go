package problem830

import (
	"reflect"
	"testing"
)

type testType struct {
	in   string
	want [][]int
}

func TestLargeGroupPositions(t *testing.T) {
	tests := [...]testType{
		{
			in: "abbxxxxzzy",
			want: [][]int{
				{3, 6},
			},
		},
		{
			in:   "abc",
			want: [][]int{},
		},
		{
			in: "abcdddeeeeaabbbcd",
			want: [][]int{
				{3, 5},
				{6, 9},
				{12, 14},
			},
		},
		{
			in: "abcddd",
			want: [][]int{
				{3, 5},
			},
		},
	}
	for _, tt := range tests {
		got := largeGroupPositions(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
