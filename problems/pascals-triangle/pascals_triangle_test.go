package problem118

import (
	"reflect"
	"testing"
)

type testType struct {
	in   int
	want [][]int
}

func TestGenerate(t *testing.T) {
	tests := [...]testType{
		{
			in: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		got := generate(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
