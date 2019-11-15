package problem696

import "testing"

type testType struct {
	in   string
	want int
}

func TestCountBinarySubstrings(t *testing.T) {
	tests := [...]testType{
		{
			in:   "00110011",
			want: 6,
		},
		{
			in:   "10101",
			want: 4,
		},
	}
	for _, tt := range tests {
		got := countBinarySubstrings(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
