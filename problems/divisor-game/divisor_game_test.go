package problem1025

import "testing"

type testType struct {
	in   int
	want bool
}

func TestDivisorGame(t *testing.T) {
	tests := [...]testType{
		{
			in:   2,
			want: true,
		},
		{
			in:   3,
			want: false,
		},
	}
	for _, tt := range tests {
		got := divisorGame(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
