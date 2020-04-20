package problem32

import "testing"

type testType struct {
	in   string
	want int
}

func TestLongestValidParentheses(t *testing.T) {
	tests := [...]testType{
		{
			in:   "(()",
			want: 2,
		},
		{
			in:   ")()())",
			want: 4,
		},
		{
			in:   ")()())(())()(",
			want: 6,
		},
		{
			in:   "()(())",
			want: 6,
		},
	}
	for _, tt := range tests {
		got := longestValidParentheses(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
