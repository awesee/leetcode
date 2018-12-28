package unique_email_addresses

import "testing"

type caseType struct {
	input    []string
	expected int
}

func TestNumUniqueEmails(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			expected: 2,
		},
		{
			input:    []string{"sandy+wang@openset.wang", "openset.wang@openset.com", "openset+wang@openset.com"},
			expected: 3,
		},
	}

	for _, tc := range tests {
		output := numUniqueEmails(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
