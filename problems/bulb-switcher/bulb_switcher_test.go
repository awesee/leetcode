package problem319

import "testing"

type testType struct {
	in   int
	want int
}

func TestBulbSwitch(t *testing.T) {
	tests := [...]testType{
		{
			in:   3,
			want: 1,
		},
		{
			in:   4,
			want: 2,
		},
		{
			in:   5,
			want: 2,
		},
	}
	for _, tt := range tests {
		got := bulbSwitch(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
