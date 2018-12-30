package long_pressed_name

import "testing"

type caseType struct {
	name     string
	typed    string
	expected bool
}

func TestIsLongPressedName(t *testing.T) {
	tests := [...]caseType{
		{
			name:     "alex",
			typed:    "aaleex",
			expected: true,
		},
		{
			name:     "saeed",
			typed:    "ssaaedd",
			expected: false,
		},
		{
			name:     "leelee",
			typed:    "lleeelee",
			expected: true,
		},
		{
			name:     "laiden",
			typed:    "laiden",
			expected: true,
		},
		{
			name:     "laiden",
			typed:    "laidef",
			expected: false,
		},
		{
			name:     "a",
			typed:    "a",
			expected: true,
		},
		{
			name:     "a",
			typed:    "b",
			expected: false,
		},
		{
			name:     "pyplrz",
			typed:    "ppyypllr",
			expected: false,
		},
		{
			name:     "vtkgn",
			typed:    "vttkgnn",
			expected: true,
		},
	}
	for _, tc := range tests {
		output := isLongPressedName(tc.name, tc.typed)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.name, tc.typed, output, tc.expected)
		}
	}
}
