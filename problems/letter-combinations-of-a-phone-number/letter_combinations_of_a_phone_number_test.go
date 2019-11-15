package problem17

import (
	"reflect"
	"testing"
)

type testType struct {
	in   string
	want []string
}

func TestLetterCombinations(t *testing.T) {
	tests := [...]testType{
		{
			in:   "",
			want: nil,
		},
		{
			in:   "23",
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}
	for _, tt := range tests {
		got := letterCombinations(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
