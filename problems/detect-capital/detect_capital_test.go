package problem520

import "testing"

type testType struct {
	in   string
	want bool
}

func TestDetectCapitalUse(t *testing.T) {
	tests := [...]testType{
		{
			in:   "USA",
			want: true,
		},
		{
			in:   "leetcode",
			want: true,
		},
		{
			in:   "Google",
			want: true,
		},
		{
			in:   "FlaG",
			want: false,
		},
	}
	for _, tt := range tests {
		got := detectCapitalUse(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
