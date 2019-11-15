package problem443

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []byte
	want []byte
}

func TestCompress(t *testing.T) {
	tests := [...]testType{
		{
			in:   []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			want: []byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			in:   []byte{'a'},
			want: []byte{'a'},
		},
		{
			in:   []byte{'a', 'a'},
			want: []byte{'a', '2'},
		},
		{
			in:   []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'},
			want: []byte{'a', 'b', '1', '2'},
		},
	}
	for _, tt := range tests {
		l := compress(tt.in)
		if !reflect.DeepEqual(tt.in[:l], tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, l, tt.want)
		}
	}
}
