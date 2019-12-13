package problem507

import "testing"

type testType struct {
	in   int
	want bool
}

func TestCheckPerfectNumber(t *testing.T) {
	tests := [...]testType{
		{
			in:   28,
			want: true,
		},
		{
			in:   6,
			want: true,
		},
		{
			in:   24,
			want: false,
		},
		{
			in:   1,
			want: false,
		},
		{
			in:   33550336,
			want: true,
		},
	}
	// Solution 1
	for _, tt := range tests {
		got := checkPerfectNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
	// Solution 2
	for _, tt := range tests {
		got := checkPerfectNumber2(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
