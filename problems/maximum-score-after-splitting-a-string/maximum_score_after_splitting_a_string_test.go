package problem1422

import "testing"

type testType struct {
	in   string
	want int
}

func TestMaxScore(t *testing.T) {
	tests := [...]testType{
		{
			in:   "011101",
			want: 5,
		},
		{
			in:   "00111",
			want: 5,
		},
		{
			in:   "1111",
			want: 3,
		},
		{
			in:   "000000",
			want: 5,
		},
	}
	for _, tt := range tests {
		got := maxScore(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
