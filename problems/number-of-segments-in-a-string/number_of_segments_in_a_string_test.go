package problem434

import "testing"

type testType struct {
	in   string
	want int
}

func TestCountSegments(t *testing.T) {
	tests := [...]testType{
		{
			in:   "Hello, my name is John",
			want: 5,
		},
		{
			in:   "",
			want: 0,
		},
		{
			in:   " ",
			want: 0,
		},
		{
			in:   " abc ",
			want: 1,
		},
		{
			in:   " abc  def  ",
			want: 2,
		},
		{
			in:   "love live! mu'sic forever",
			want: 4,
		},
	}
	for _, tt := range tests {
		got := countSegments(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
