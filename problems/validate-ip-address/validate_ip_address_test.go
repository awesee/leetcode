package problem_468

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestValidIPAddress(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "172.16.254.1",
			expected: "IPv4",
		},
		{
			input:    "172.16.254.01",
			expected: "Neither",
		},
		{
			input:    "256.256.256.256",
			expected: "Neither",
		},
		{
			input:    "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: "IPv6",
		},
		{
			input:    "2001:db8:85a3:0:0:8A2E:0370:7334",
			expected: "IPv6",
		},
		{
			input:    "2001:0db8:85a3::8A2E:0370:7334",
			expected: "Neither",
		},
		{
			input:    "02001:0db8:85a3:0000:0000:8a2e:0370:7334",
			expected: "Neither",
		},
	}
	for _, tc := range tests {
		output := validIPAddress(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
