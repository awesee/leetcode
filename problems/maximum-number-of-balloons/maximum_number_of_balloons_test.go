package problem1189

import "testing"

type testType struct {
	in   string
	want int
}

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := [...]testType{
		{
			in:   "nlaebolko",
			want: 1,
		},
		{
			in:   "loonbalxballpoon",
			want: 2,
		},
		{
			in:   "leetcode",
			want: 0,
		},
		{
			in:   "lloo",
			want: 0,
		},
	}
	for _, tt := range tests {
		got := maxNumberOfBalloons(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
