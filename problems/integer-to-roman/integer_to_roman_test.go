package problem12

import "testing"

type testType struct {
	in   int
	want string
}

func TestIntToRoman(t *testing.T) {
	tests := [...]testType{
		{
			in:   3,
			want: "III",
		},
		{
			in:   4,
			want: "IV",
		},
		{
			in:   9,
			want: "IX",
		},
		{
			in:   20,
			want: "XX",
		},
		{
			in:   58,
			want: "LVIII",
		},
		{
			in:   1994,
			want: "MCMXCIV",
		},
		{
			in:   11111,
			want: "MMMMMMMMMMMCXI",
		},
	}
	for _, tt := range tests {
		got := intToRoman(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intToRoman(111111)
	}
}
