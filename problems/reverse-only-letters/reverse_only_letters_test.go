package reverse_only_letters

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestReverseOnlyLetters(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "ab-cd",
			expected: "dc-ba",
		},
		{
			input:    "a-bC-dEf-ghIj",
			expected: "j-Ih-gfE-dCba",
		},
		{
			input:    "Test1ng-Leet=code-Q!",
			expected: "Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tc := range tests {
		output := reverseOnlyLetters(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
