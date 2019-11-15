package problem541

import "testing"

type testType struct {
	s    string
	k    int
	want string
}

func TestReverseStr(t *testing.T) {
	tests := [...]testType{
		{
			s:    "abcdefg",
			k:    2,
			want: "bacdfeg",
		},
		{
			s:    "abcdefg",
			k:    4,
			want: "dcbaefg",
		},
		{
			s:    "abcd",
			k:    4,
			want: "dcba",
		},
		{
			s:    "abcdefg",
			k:    8,
			want: "gfedcba",
		},
	}
	for _, tt := range tests {
		got := reverseStr(tt.s, tt.k)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.s, tt.k, got, tt.want)
		}
	}
}
