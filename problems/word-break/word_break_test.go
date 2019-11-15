package problem139

import "testing"

type testType struct {
	in       string
	wordDict []string
	want     bool
}

func TestWordBreak(t *testing.T) {
	tests := [...]testType{
		{
			in:       "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			in:       "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			in:       "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
	}
	for _, tt := range tests {
		got := wordBreak(tt.in, tt.wordDict)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
