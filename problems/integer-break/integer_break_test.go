package problem343

import "testing"

type testType struct {
	in   int
	want int
}

func TestIntegerBreak(t *testing.T) {
	tests := [...]testType{
		{
			in:   2,
			want: 1,
		},
		{
			in:   10,
			want: 36,
		},
		{
			in:   3,
			want: 2,
		},
		{
			in:   7,
			want: 12,
		},
		{
			in:   17,
			want: 486,
		},
	}
	for _, tt := range tests {
		got := integerBreak(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
