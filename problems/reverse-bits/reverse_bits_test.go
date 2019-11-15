package problem190

import "testing"

type testType struct {
	in   uint32
	want uint32
}

func TestReverseBits(t *testing.T) {
	tests := [...]testType{
		{
			in:   43261596,
			want: 964176192,
		},
		{
			in:   4294967293,
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		got := reverseBits(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
