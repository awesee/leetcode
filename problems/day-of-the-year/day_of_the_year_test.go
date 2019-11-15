package problem1154

import "testing"

type testType struct {
	in   string
	want int
}

func TestDayOfYear(t *testing.T) {
	tests := [...]testType{
		{
			in:   "2019-01-09",
			want: 9,
		},
		{
			in:   "2019-02-10",
			want: 41,
		},
		{
			in:   "2003-03-01",
			want: 60,
		},
		{
			in:   "2004-03-01",
			want: 61,
		},
		{
			in:   "2000-08-01",
			want: 214,
		},
		{
			in:   "1993-12-11",
			want: 345,
		},
	}
	for _, tt := range tests {
		got := dayOfYear(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
