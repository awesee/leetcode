package problem521

import "testing"

type testType struct {
	a    string
	b    string
	want int
}

func TestFindLUSlength(t *testing.T) {
	tests := [...]testType{
		{
			a:    "aba",
			b:    "cdc",
			want: 3,
		},
		{
			a:    "aba",
			b:    "abcdef",
			want: 6,
		},
		{
			a:    "abcdef",
			b:    "cdc",
			want: 6,
		},
		{
			a:    "abc",
			b:    "abc",
			want: -1,
		},
	}
	for _, tt := range tests {
		got := findLUSlength(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
