package problem1009

import "testing"

type testType struct {
	in   int
	want int
}

func TestBitwiseComplement(t *testing.T) {
	tests := [...]testType{
		{
			in:   5,
			want: 2,
		},
		{
			in:   7,
			want: 0,
		},
		{
			in:   10,
			want: 5,
		},
		{
			in:   0,
			want: 1,
		},
	}
	for _, tt := range tests {
		got := bitwiseComplement(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
