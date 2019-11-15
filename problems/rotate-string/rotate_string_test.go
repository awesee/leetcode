package problem796

import "testing"

type testType struct {
	A    string
	B    string
	want bool
}

func TestRotateString(t *testing.T) {
	tests := [...]testType{
		{
			A:    "abcde",
			B:    "cdeab",
			want: true,
		},
		{
			A:    "abcde",
			B:    "abced",
			want: false,
		},
		{
			A:    "abcde",
			B:    "abcdef",
			want: false,
		},
	}
	for _, tt := range tests {
		got := rotateString(tt.A, tt.B)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.A, tt.B, got, tt.want)
		}
	}
}
