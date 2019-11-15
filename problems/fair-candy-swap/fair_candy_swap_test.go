package problem888

import (
	"reflect"
	"testing"
)

type testType struct {
	a    []int
	b    []int
	want []int
}

func TestFairCandySwap(t *testing.T) {
	tests := [...]testType{
		{
			a:    []int{1, 1},
			b:    []int{2, 2},
			want: []int{1, 2},
		},
		{
			a:    []int{1, 2},
			b:    []int{2, 3},
			want: []int{1, 2},
		},
		{
			a:    []int{2},
			b:    []int{1, 3},
			want: []int{2, 3},
		},
		{
			a:    []int{1, 2, 5},
			b:    []int{2, 4},
			want: []int{5, 4},
		},
		{
			a:    []int{1, 2, 3},
			b:    []int{11, 15, 17},
			want: nil,
		},
	}
	for _, tt := range tests {
		got := fairCandySwap(tt.a, tt.b)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
