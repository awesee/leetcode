package problem171

import "testing"

type testType struct {
	in   string
	want int
}

func TestTitleToNumber(t *testing.T) {
	tests := [...]testType{
		{
			in:   "A",
			want: 1,
		},
		{
			in:   "C",
			want: 3,
		},
		{
			in:   "Z",
			want: 26,
		},
		{
			in:   "AA",
			want: 27,
		},
		{
			in:   "AB",
			want: 28,
		},
		{
			in:   "ZY",
			want: 701,
		}, {
			in:   "AAA",
			want: 703,
		},
		{
			in:   "XYZ",
			want: 16900,
		},
	}
	for _, tt := range tests {
		got := titleToNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
