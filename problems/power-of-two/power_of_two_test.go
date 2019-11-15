package problem231

import "testing"

type testType struct {
	in   int
	want bool
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := [...]testType{
		{
			in:   0,
			want: false,
		},
		{
			in:   1,
			want: true,
		},
		{
			in:   12,
			want: false,
		},
		{
			in:   16,
			want: true,
		},
		{
			in:   218,
			want: false,
		},
	}
	for _, tt := range tests {
		got := isPowerOfTwo(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
