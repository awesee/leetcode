package problem788

import "testing"

type testType struct {
	in   int
	want int
}

func TestRotatedDigits(t *testing.T) {
	tests := [...]testType{
		{
			in:   10,
			want: 4,
		},
		{
			in:   23,
			want: 11,
		},
		{
			in:   100,
			want: 40,
		},
		{
			in:   200,
			want: 81,
		},
	}
	for _, tt := range tests {
		got := rotatedDigits(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
