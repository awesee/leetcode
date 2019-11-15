package problem345

import "testing"

type testType struct {
	in   string
	want string
}

func TestReverseVowels(t *testing.T) {
	tests := [...]testType{
		{
			in:   "hello",
			want: "holle",
		},
		{
			in:   "leetcode",
			want: "leotcede",
		},
	}
	for _, tt := range tests {
		got := reverseVowels(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
