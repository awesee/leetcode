package fizz_buzz

import "testing"

type caseType struct {
	input    int
	expected []string
}

func TestFizzBuzz(t *testing.T) {
	tests := [...]caseType{
		{
			input: 15,
			expected: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
		{
			input: 5,
			expected: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
			},
		},
		{
			input: 10,
			expected: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
			},
		},
	}

	for _, tc := range tests {
		output := fizzBuzz(tc.input)
		if len(output) != len(tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
		for k, v := range tc.expected {
			if output[k] != v {
				t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
			}
		}
	}
}
