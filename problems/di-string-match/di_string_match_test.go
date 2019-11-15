package problem942

import (
	"reflect"
	"testing"
)

type testType struct {
	in   string
	want []int
}

func TestDiStringMatch(t *testing.T) {
	tests := [...]testType{
		{
			in:   "IDID",
			want: []int{0, 4, 1, 3, 2},
		},
		{
			in:   "III",
			want: []int{0, 1, 2, 3},
		},
		{
			in:   "DDI",
			want: []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		got := diStringMatch(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
