package problem202

import "testing"

type testType struct {
	in   int
	want bool
}

func TestIsHappy(t *testing.T) {
	tests := [...]testType{
		{
			in:   19,
			want: true,
		},
		{
			in:   1,
			want: true,
		},
		{
			in:   0,
			want: false,
		},
	}
	for _, tt := range tests {
		got := isHappy(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
