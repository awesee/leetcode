package problem65

import "testing"

type testType struct {
	in   string
	want bool
}

func TestIsNumber(t *testing.T) {
	tests := [...]testType{
		{
			in:   "0",
			want: true,
		},
		{
			in:   " 0.1",
			want: true,
		},
		{
			in:   "abc",
			want: false,
		},
		{
			in:   "1 a",
			want: false,
		},
		{
			in:   "2e10",
			want: true,
		},
		{
			in:   " -90e3   ",
			want: true,
		},
		{
			in:   " 1e",
			want: false,
		},
		{
			in:   "e3",
			want: false,
		},
		{
			in:   " 6e-1",
			want: true,
		},
		{
			in:   " 99e2.5 ",
			want: false,
		},
		{
			in:   "53.5e93",
			want: true,
		},
		{
			in:   " --6 ",
			want: false,
		},
		{
			in:   "-+3",
			want: false,
		},
		{
			in:   "95a54e53",
			want: false,
		},
		{
			in:   "078332e437",
			want: true,
		},
	}
	for _, tt := range tests {
		got := isNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
