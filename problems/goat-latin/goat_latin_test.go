package goat_latin

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestToGoatLatin(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "I speak Goat Latin",
			expected: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			input:    "The quick brown fox jumped over the lazy dog",
			expected: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tc := range tests {
		output := toGoatLatin(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
