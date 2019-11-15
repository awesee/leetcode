package problem3

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
		"abcdef":   6,
	}

	for in, want := range tests {
		got := lengthOfLongestSubstring(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
