package problem5

import "testing"

type testType struct {
	in   string
	want string
}

func TestLongestPalindrome(t *testing.T) {
	tests := [...]testType{
		{
			in:   "babad",
			want: "bab",
		},
		{
			in:   "cbbd",
			want: "bb",
		},
		{
			in:   "",
			want: "",
		},
		{
			in:   "a",
			want: "a",
		},
		{
			in:   "abbba",
			want: "abbba",
		},
	}
	for _, tt := range tests {
		got := longestPalindrome(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
