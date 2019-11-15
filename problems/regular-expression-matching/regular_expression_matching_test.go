package problem10

import "testing"

type testType struct {
	s    string
	p    string
	want bool
}

func TestIsMatch(t *testing.T) {
	tests := [...]testType{
		{
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			s:    "mississippi",
			p:    "mis*is*p*.",
			want: false,
		},
		{
			s:    "",
			p:    ".",
			want: false,
		},
	}
	for _, tt := range tests {
		got := isMatch(tt.s, tt.p)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.s, tt.p, got, tt.want)
		}
	}
}
