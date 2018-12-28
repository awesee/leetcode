package to_lower_case

import "testing"

func TestToLowerCase(t *testing.T) {
	tests := map[string]string{
		"Hello":   "hello",
		"here":    "here",
		"LOVELY":  "lovely",
		"OpenSet": "openset",
	}

	for input, expected := range tests {
		output := toLowerCase(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
