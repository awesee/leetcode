package problem387

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := map[string]int{
		"leetcode":     0,
		"loveleetcode": 2,
		"hello world":  0,
		"aabbccddeeff": -1,
	}

	for in, want := range tests {
		got := firstUniqChar(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
