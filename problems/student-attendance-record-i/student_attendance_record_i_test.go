package problem551

import "testing"

type testType struct {
	in   string
	want bool
}

func TestCheckRecord(t *testing.T) {
	tests := [...]testType{
		{
			in:   "PPALLP",
			want: true,
		},
		{
			in:   "PPALLL",
			want: false,
		},
	}
	for _, tt := range tests {
		got := checkRecord(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
