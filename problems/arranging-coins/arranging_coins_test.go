package problem441

import "testing"

type testType struct {
	in   int
	want int
}

func TestArrangeCoins(t *testing.T) {
	tests := [...]testType{
		{
			in:   5,
			want: 2,
		},
		{
			in:   8,
			want: 3,
		},
		{
			in:   0,
			want: 0,
		},
		{
			in:   1,
			want: 1,
		},
		{
			in:   2,
			want: 1,
		},
		{
			in:   3,
			want: 2,
		},
		{
			in:   13,
			want: 4,
		},
		{
			in:   130,
			want: 15,
		},
	}
	for _, tt := range tests {
		got := arrangeCoins(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
