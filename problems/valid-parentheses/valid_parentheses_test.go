package problem20

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

	for in, want := range tests {
		got := isValid(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValid("{[({[({[()]})]})]}")
	}
}
