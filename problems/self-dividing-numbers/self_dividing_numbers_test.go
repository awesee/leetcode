package problem728

import (
	"reflect"
	"testing"
)

type testType struct {
	left  int
	right int
	want  []int
}

func TestSelfDividingNumbers(t *testing.T) {
	tests := [...]testType{
		{
			left:  1,
			right: 22,
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
	}
	for _, tt := range tests {
		got := selfDividingNumbers(tt.left, tt.right)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.left, tt.right, got, tt.want)
		}
	}
}
