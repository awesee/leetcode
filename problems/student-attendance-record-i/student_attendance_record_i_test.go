package student_attendance_record_i

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestCheckRecord(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "PPALLP",
			expected: true,
		},
		{
			input:    "PPALLL",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := checkRecord(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
