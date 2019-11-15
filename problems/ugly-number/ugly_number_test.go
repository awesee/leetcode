package problem263

import "testing"

type testType struct {
	in   int
	want bool
}

func TestIsUgly(t *testing.T) {
	tests := [...]testType{
		{
			in:   1,
			want: true,
		},
		{
			in:   6,
			want: true,
		},
		{
			in:   8,
			want: true,
		},
		{
			in:   14,
			want: false,
		},
		{
			in:   0,
			want: false,
		},
		{
			in:   -30,
			want: false,
		},
	}
	for _, tt := range tests {
		got := isUgly(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
