package problem744

import "testing"

type testType struct {
	in     []byte
	target byte
	want   byte
}

func TestNextGreatestLetter(t *testing.T) {
	tests := [...]testType{
		{
			in:     []byte{'c', 'f', 'j'},
			target: 'a',
			want:   'c',
		},
		{
			in:     []byte{'c', 'f', 'j'},
			target: 'c',
			want:   'f',
		},
		{
			in:     []byte{'c', 'f', 'j'},
			target: 'd',
			want:   'f',
		},
		{
			in:     []byte{'c', 'f', 'j'},
			target: 'g',
			want:   'j',
		},
		{
			in:     []byte{'c', 'f', 'j'},
			target: 'j',
			want:   'c',
		},
		{
			in:     []byte{'c', 'f', 'j'},
			target: 'k',
			want:   'c',
		},
	}
	for _, tt := range tests {
		got := nextGreatestLetter(tt.in, tt.target)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
