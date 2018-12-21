package first_unique_character_in_a_string

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := map[string]int{
		"leetcode":     0,
		"loveleetcode": 2,
		"hello world":  0,
		"aabbccddeeff": -1,
	}

	for input, expected := range tests {
		output := firstUniqChar(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
