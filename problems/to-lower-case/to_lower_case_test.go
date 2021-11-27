package problem709

import "testing"

func TestToLowerCase(t *testing.T) {
	tests := map[string]string{
		"Hello":  "hello",
		"here":   "here",
		"LOVELY": "lovely",
		"aweSee": "awesee",
	}

	for in, want := range tests {
		got := toLowerCase(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
