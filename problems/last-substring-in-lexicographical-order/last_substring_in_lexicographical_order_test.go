package problem1163

import "testing"

type testType struct {
	in   string
	want string
}

func TestLastSubstring(t *testing.T) {
	tests := [...]testType{
		{
			in:   "abab",
			want: "bab",
		},
		{
			in:   "abcdabc",
			want: "dabc",
		},
		{
			in:   "ffbcdfeda",
			want: "ffbcdfeda",
		},
		{
			in:   "hksljkjsfjv",
			want: "v",
		},
	}
	for _, tt := range tests {
		got := lastSubstring(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
