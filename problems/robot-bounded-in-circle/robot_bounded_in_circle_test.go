package problem1041

import "testing"

type testType struct {
	in   string
	want bool
}

func TestIsRobotBounded(t *testing.T) {
	tests := [...]testType{
		{
			in:   "GGLLGG",
			want: true,
		},
		{
			in:   "GG",
			want: false,
		},
		{
			in:   "GL",
			want: true,
		},
		{
			in:   "GGLLGGGGRRGG",
			want: true,
		},
		{
			in:   "GGRRGG",
			want: true,
		},
		{
			in:   "GLRLLGLL",
			want: true,
		},
		{
			in:   "GLGRGLGLGLGL",
			want: false,
		},
	}
	for _, tt := range tests {
		got := isRobotBounded(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
