package problem326

import "testing"

type testType struct {
	in   int
	want bool
}

func TestIsPowerOfThree(t *testing.T) {
	tests := [...]testType{
		{
			in:   27,
			want: true,
		},
		{
			in:   0,
			want: false,
		},
		{
			in:   1,
			want: true,
		},
		{
			in:   9,
			want: true,
		},
		{
			in:   45,
			want: false,
		},
	}
	for _, tt := range tests {
		got := isPowerOfThree(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
