package problem155

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestMinStack(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{-2, 0, -3},
			want: []int{-3, 0, -2},
		},
	}
	for _, tt := range tests {
		minStack, got := Constructor(), make([]int, 0)
		for _, x := range tt.in {
			minStack.Push(x)
		}
		got = append(got, minStack.GetMin())
		minStack.Pop()
		got = append(got, minStack.Top())
		got = append(got, minStack.GetMin())
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
