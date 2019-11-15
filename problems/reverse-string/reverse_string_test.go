package problem344

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []byte
	want []byte
}

func TestReverseString(t *testing.T) {
	tests := [...]testType{
		{
			in:   []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			in:   []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		got := make([]byte, len(tt.in))
		copy(got, tt.in)
		reverseString(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
