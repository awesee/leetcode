package problem1002

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []string
	want []string
}

func TestCommonChars(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"bella", "label", "roller"},
			want: []string{"e", "l", "l"},
		},
		{
			in:   []string{"cool", "lock", "cook"},
			want: []string{"c", "o"},
		},
	}
	for _, tt := range tests {
		got := commonChars(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
