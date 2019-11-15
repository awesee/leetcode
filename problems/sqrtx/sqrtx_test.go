package problem69

import "testing"

type testType struct {
	in   int
	want int
}

func TestMySqrt(t *testing.T) {
	tests := [...]testType{
		{
			in:   4,
			want: 2,
		},
		{
			in:   8,
			want: 2,
		},
		{
			in:   0,
			want: 0,
		},
		{
			in:   1,
			want: 1,
		},
	}
	for _, tt := range tests {
		got := mySqrt(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
