package problem893

import "testing"

type testType struct {
	in   []string
	want int
}

func TestNumSpecialEquivGroups(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"a", "b", "c", "a", "c", "c"},
			want: 3,
		},
		{
			in:   []string{"aa", "bb", "ab", "ba"},
			want: 4,
		},
		{
			in:   []string{"abc", "acb", "bac", "bca", "cab", "cba"},
			want: 3,
		},
		{
			in:   []string{"abcd", "cdab", "adcb", "cbad"},
			want: 1,
		},
	}
	for _, tt := range tests {
		got := numSpecialEquivGroups(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
