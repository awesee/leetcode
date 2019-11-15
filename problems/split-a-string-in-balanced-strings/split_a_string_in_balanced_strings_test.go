package problem1221

import "testing"

type testType struct {
	in   string
	want int
}

func TestBalancedStringSplit(t *testing.T) {
	tests := [...]testType{
		{
			in:   "RLRRLLRLRL",
			want: 4,
		},
		{
			in:   "RLLLLRRRLR",
			want: 3,
		},
		{
			in:   "LLLLRRRR",
			want: 1,
		},
	}
	for _, tt := range tests {
		got := balancedStringSplit(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
