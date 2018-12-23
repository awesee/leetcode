package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	tests := map[string]string{
		"hello, world":                   "dlrow ,olleh",
		"你好，世界":                          "界世，好你",
		"A man, a plan, a canal: Panama": "amanaP :lanac a ,nalp a ,nam A",
	}

	for input, expected := range tests {
		output := reverseString(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
