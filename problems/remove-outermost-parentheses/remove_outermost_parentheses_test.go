package problem1021

import "testing"

type testType struct {
	in   string
	want string
}

func TestRemoveOuterParentheses(t *testing.T) {
	tests := [...]testType{
		{
			in:   "(()())(())",
			want: "()()()",
		},
		{
			in:   "(()())(())(()(()))",
			want: "()()()()(())",
		},
		{
			in:   "()()",
			want: "",
		},
	}
	for _, tt := range tests {
		got := removeOuterParentheses(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
