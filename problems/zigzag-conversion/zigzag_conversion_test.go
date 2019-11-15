package problem6

import "testing"

type testType struct {
	in      string
	numRows int
	want    string
}

func TestConvert(t *testing.T) {
	tests := [...]testType{
		{
			in:      "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			in:      "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			in:      "AB",
			numRows: 1,
			want:    "AB",
		},
	}
	for _, tt := range tests {
		got := convert(tt.in, tt.numRows)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
