package problem172

import "testing"

type testType struct {
	in   int
	want int
}

func TestTrailingZeroes(t *testing.T) {
	tests := [...]testType{
		{
			in:   3,
			want: 0,
		},
		{
			in:   5,
			want: 1,
		},
		{
			in:   7,
			want: 1,
		},
		{
			in:   8,
			want: 1,
		},
		{
			in:   10,
			want: 2,
		},
		{
			in:   12,
			want: 2,
		},
		{
			in:   25,
			want: 6,
		},
	}
	for _, tt := range tests {
		got := trailingZeroes(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
