package problem151

import "testing"

type testType struct {
	in   string
	want string
}

func TestReverseWords(t *testing.T) {
	tests := [...]testType{
		{
			in:   "the sky is blue",
			want: "blue is sky the",
		},
		{
			in:   "  hello world!  ",
			want: "world! hello",
		},
		{
			in:   "a good   example",
			want: "example good a",
		},
	}
	for _, tt := range tests {
		got := reverseWords(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
