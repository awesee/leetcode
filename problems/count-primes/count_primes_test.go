package problem204

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

	for in, want := range tests {
		got := countPrimes(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
