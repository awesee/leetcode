package problem680

import "testing"

type testType struct {
	in   string
	want bool
}

func TestValidPalindrome(t *testing.T) {
	tests := [...]testType{
		{
			in:   "aba",
			want: true,
		},
		{
			in:   "abca",
			want: true,
		},
		{
			in:   "hello",
			want: false,
		},
		{
			in:   "abcdcbda",
			want: true,
		},
		{
			in:   "abcbabbbca",
			want: false,
		},
	}
	for _, tt := range tests {
		got := validPalindrome(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
