package problem119

import (
	"reflect"
	"testing"
)

type testType struct {
	in   int
	want []int
}

func TestGetRow(t *testing.T) {
	tests := [...]testType{
		{
			in:   2,
			want: []int{1, 2, 1},
		},
		{
			in:   3,
			want: []int{1, 3, 3, 1},
		},
		{
			in:   5,
			want: []int{1, 5, 10, 10, 5, 1},
		},
	}
	for _, tt := range tests {
		got := getRow(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
