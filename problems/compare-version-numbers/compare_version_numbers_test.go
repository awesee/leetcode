package compare_version_numbers

import "testing"

type caseType struct {
	v1       string
	v2       string
	expected int
}

func TestCompareVersion(t *testing.T) {
	tests := [...]caseType{
		{
			v1:       "1.0.01",
			v2:       "1.00.1",
			expected: 0,
		},
		{
			v1:       "1.0.0",
			v2:       "1",
			expected: 0,
		},
		{
			v1:       "1.01",
			v2:       "1.001",
			expected: 0,
		},
		{
			v1:       "1.0",
			v2:       "1.0.0",
			expected: 0,
		},
		{
			v1:       "0.1",
			v2:       "1.1",
			expected: -1,
		},
		{
			v1:       "7.5.2.4",
			v2:       "7.5.3",
			expected: -1,
		},

		{
			v1:       "7.00",
			v2:       "7.001",
			expected: -1,
		},
		{
			v1:       "1.1",
			v2:       "1.10",
			expected: -1,
		},
		{
			v1:       "1.0.1",
			v2:       "1",
			expected: 1,
		},
		{
			v1:       "1.10",
			v2:       "1.2",
			expected: 1,
		},
		{
			v1:       "1.10.05",
			v2:       "1.9.10",
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := compareVersion(tc.v1, tc.v2)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.v1, tc.v2, output, tc.expected)
		}
	}
}
