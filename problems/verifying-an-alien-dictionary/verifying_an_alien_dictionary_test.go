package problem953

import "testing"

type testType struct {
	words []string
	order string
	want  bool
}

func TestIsAlienSorted(t *testing.T) {
	tests := [...]testType{
		{
			words: []string{"hello", "leetcode"},
			order: "hlabcdefgijkmnopqrstuvwxyz",
			want:  true,
		},
		{
			words: []string{"word", "world", "row"},
			order: "worldabcefghijkmnpqstuvxyz",
			want:  false,
		},
		{
			words: []string{"apple", "app"},
			order: "abcdefghijklmnopqrstuvwxyz",
			want:  false,
		},
	}
	for _, tt := range tests {
		got := isAlienSorted(tt.words, tt.order)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.words, got, tt.want)
		}
	}
}
