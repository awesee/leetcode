package problem14

import "testing"

type testType struct {
	in   []string
	want string
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			in:   []string{"dog", "race", "car"},
			want: "",
		},
		{
			in:   nil,
			want: "",
		},
	}

	for _, tt := range tests {
		got := longestCommonPrefix(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
