package word_break

import "testing"

type caseType struct {
	input    string
	wordDict []string
	expected bool
}

func TestWordBreak(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			input:    "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			input:    "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := wordBreak(tc.input, tc.wordDict)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
