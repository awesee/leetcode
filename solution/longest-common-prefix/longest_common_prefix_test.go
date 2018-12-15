package longest_common_prefix

import "testing"

type caseType struct {
	Input    []string
	Expected string
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []caseType{
		{
			Input:    []string{"flower", "flow", "flight"},
			Expected: "fl",
		},
		{
			Input:    []string{"dog", "race", "car"},
			Expected: "",
		},
		{
			Input:    nil,
			Expected: "",
		},
	}

	for _, test := range tests {
		output := longestCommonPrefix(test.Input)
		if output != test.Expected {
			t.Fatalf("Input: %v, Output: %v, Expected: %v", test.Input, output, test.Expected)
		}
	}
}
