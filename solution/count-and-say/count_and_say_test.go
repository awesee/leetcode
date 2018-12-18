package count_and_say

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := map[int]string{
		1: "1",
		3: "21",
		5: "111221",
		7: "13112221",
		9: "31131211131221",
	}

	for input, expected := range tests {
		output := countAndSay(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
