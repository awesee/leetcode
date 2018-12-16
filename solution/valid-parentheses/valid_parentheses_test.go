package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := map[string]bool{
		"":       true,
		"()":     true,
		"()[]{}": true,
		"{[]}":   true,
		"(]":     false,
		"([)]":   false,
		"([{)]":  false,
		"([})]":  false,
	}

	for input, expected := range tests {
		output := isValid(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValid("{[({[({[()]})]})]}")
	}
}
