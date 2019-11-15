package problem342

import "testing"

type testType struct {
	in   int
	want bool
}

func TestIsPowerOfFour(t *testing.T) {
	tests := [...]testType{
		{
			in:   16,
			want: true,
		},
		{
			in:   5,
			want: false,
		},
		{
			in:   6,
			want: false,
		},
		{
			in:   12,
			want: false,
		},
		{
			in:   0,
			want: false,
		},
		{
			in:   1,
			want: true,
		},
	}
	for _, tt := range tests {
		got := isPowerOfFour(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
