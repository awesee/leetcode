package problem292

import "testing"

type testType struct {
	in   int
	want bool
}

func TestCanWinNim(t *testing.T) {
	tests := [...]testType{
		{
			in:   4,
			want: false,
		},
		{
			in:   3,
			want: true,
		},
	}
	for _, tt := range tests {
		got := canWinNim(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
