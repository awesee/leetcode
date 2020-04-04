package problem1332

import "testing"

type testType struct {
	in   string
	want int
}

func TestRemovePalindromeSub(t *testing.T) {
	tests := [...]testType{
		{
			in:   "ababa",
			want: 1,
		},
		{
			in:   "abb",
			want: 2,
		},
		{
			in:   "baabb",
			want: 2,
		},
		{
			in:   "",
			want: 0,
		},
		{
			in:   "aaabbbaaabbb",
			want: 2,
		},
	}
	for _, tt := range tests {
		got := removePalindromeSub(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
