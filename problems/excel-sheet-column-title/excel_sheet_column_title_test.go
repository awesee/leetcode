package problem168

import "testing"

type testType struct {
	in   int
	want string
}

func TestConvertToTitle(t *testing.T) {
	tests := [...]testType{
		{
			in:   1,
			want: "A",
		},
		{
			in:   3,
			want: "C",
		},
		{
			in:   26,
			want: "Z",
		},
		{
			in:   52,
			want: "AZ",
		},
		{
			in:   27,
			want: "AA",
		},
		{
			in:   28,
			want: "AB",
		},
		{
			in:   701,
			want: "ZY",
		},
		{
			in:   703,
			want: "AAA",
		},
		{
			in:   16900,
			want: "XYZ",
		},
	}
	// Solution 1
	for _, tt := range tests {
		got := convertToTitle(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
	// Solution 2
	for _, tt := range tests {
		got := convertToTitle2(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
