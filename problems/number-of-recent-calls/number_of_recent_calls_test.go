package problem933

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []int
	want []int
}

func TestConstructor(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 100, 3001, 3002},
			want: []int{1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		obj, got := Constructor(), make([]int, 0)
		for _, t := range tt.in {
			got = append(got, obj.Ping(t))
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
