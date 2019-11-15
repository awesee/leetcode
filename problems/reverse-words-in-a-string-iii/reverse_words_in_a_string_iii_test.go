package problem557

import "testing"

type testType struct {
	in   string
	want string
}

func TestReverseWords(t *testing.T) {
	tests := [...]testType{
		{
			in:   "Let's take LeetCode contest",
			want: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		got := reverseWords(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
