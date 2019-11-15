package problem264

import "testing"

type testType struct {
	in   int
	want int
}

func TestNthUglyNumber(t *testing.T) {
	tests := [...]testType{
		{
			in:   10,
			want: 12,
		},
		{
			in:   9,
			want: 10,
		},
		{
			in:   60,
			want: 384,
		},
		{
			in:   90,
			want: 1152,
		},
		{
			in:   1,
			want: 1,
		},
	}
	for _, tt := range tests {
		got := nthUglyNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
