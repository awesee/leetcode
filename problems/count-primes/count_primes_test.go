package count_primes

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := map[int]int{
		0:   0,
		1:   0,
		2:   0,
		3:   1,
		5:   2,
		7:   3,
		10:  4,
		100: 25,
		200: 46,
		300: 62,
	}

	for input, expected := range tests {
		output := countPrimes(input)
		if output != expected {
			t.Fatalf("input: %v, output: %v, expected: %v", input, output, expected)
		}
	}
}
