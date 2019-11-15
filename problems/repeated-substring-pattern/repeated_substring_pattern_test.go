package problem459

import "testing"

type testType struct {
	in   string
	want bool
}

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := [...]testType{
		{
			in:   "a",
			want: false,
		},
		{
			in:   "abab",
			want: true,
		},
		{
			in:   "aba",
			want: false,
		},
		{
			in:   "abaaba",
			want: true,
		},
		{
			in:   "abcabcabcabc",
			want: true,
		},
		{
			in:   "abbaabbaabba",
			want: true,
		},
		{
			in:   "abcabcabcdabcabcabcdabcabcabcd",
			want: true,
		},
		{
			in:   "abaabaabac",
			want: false,
		}, {
			in:   "babbabbabbabbab",
			want: true,
		},
	}
	for _, tt := range tests {
		got := repeatedSubstringPattern(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
