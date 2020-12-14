package problem1662

import "testing"

type testType struct {
	word1 []string
	word2 []string
	want  bool
}

func TestArrayStringsAreEqual(t *testing.T) {
	tests := [...]testType{
		{
			word1: []string{"ab", "c"},
			word2: []string{"a", "bc"},
			want:  true,
		},
		{
			word1: []string{"a", "cb"},
			word2: []string{"ab", "c"},
			want:  false,
		},
		{
			word1: []string{"abc", "d", "defg"},
			word2: []string{"abcddefg"},
			want:  true,
		},
	}
	for _, tt := range tests {
		got := arrayStringsAreEqual(tt.word1, tt.word2)
		if got != tt.want {
			t.Fatalf("word1: %v, word2: %v, got: %v, want: %v", tt.word1, tt.word2, got, tt.want)
		}
	}
}
