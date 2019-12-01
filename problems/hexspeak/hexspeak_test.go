package problem1271

import "testing"

type testType struct {
	in   string
	want string
}

func TestToHexspeak(t *testing.T) {
	tests := [...]testType{
		{
			in:   "257",
			want: "IOI",
		},
		{
			in:   "3",
			want: "ERROR",
		},
	}
	for _, tt := range tests {
		got := toHexspeak(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
