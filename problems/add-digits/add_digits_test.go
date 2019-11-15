package problem258

import "testing"

type testType struct {
	in   int
	want int
}

func TestAddDigits(t *testing.T) {
	tests := [...]testType{
		{
			in:   38,
			want: 2,
		},
	}
	for _, tt := range tests {
		got := addDigits(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
