package problem1408

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []string
	want []string
}

func TestStringMatching(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"mass", "as", "hero", "superhero"},
			want: []string{"as", "hero"},
		},
		{
			in:   []string{"leetcode", "et", "code"},
			want: []string{"et", "code"},
		},
		{
			in:   []string{"blue", "green", "bu"},
			want: nil,
		},
	}
	for _, tt := range tests {
		got := stringMatching(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
