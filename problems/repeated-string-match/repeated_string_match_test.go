package problem686

import "testing"

type testType struct {
	a    string
	b    string
	want int
}

func TestRepeatedStringMatch(t *testing.T) {
	tests := [...]testType{
		{
			a:    "abcd",
			b:    "cdabcdab",
			want: 3,
		},
		{
			a:    "abcdefg",
			b:    "cde",
			want: 1,
		},
		{
			a:    "abc",
			b:    "abcabc",
			want: 2,
		},
		{
			a:    "abc",
			b:    "abcd",
			want: -1,
		},
	}
	for _, tt := range tests {
		got := repeatedStringMatch(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
