package problem233

import "testing"

type testType struct {
	in   int
	want int
}

func TestCountDigitOne(t *testing.T) {
	tests := [...]testType{
		{
			in:   0,
			want: 0,
		},
		{
			in:   1,
			want: 1,
		},
		{
			in:   12,
			want: 5,
		},
		{
			in:   99,
			want: 20,
		},
		{
			in:   100,
			want: 21,
		},
		{
			in:   123,
			want: 57,
		},
		{
			in:   1234,
			want: 689,
		},
		{
			in:   12345,
			want: 8121,
		},
		{
			in:   123456,
			want: 93553,
		},
		{
			in:   1234567,
			want: 1058985,
		},
		{
			in:   12345678,
			want: 11824417,
		},
		{
			in:   123456789,
			want: 130589849,
		},
	}
	for _, tt := range tests {
		got := countDigitOne(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
