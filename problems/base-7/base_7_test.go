package problem504

import "testing"

type testType struct {
	in   int
	want string
}

func TestConvertToBase7(t *testing.T) {
	tests := [...]testType{
		{
			in:   100,
			want: "202",
		},
		{
			in:   -7,
			want: "-10",
		},
		{
			in:   -1,
			want: "-1",
		},
		{
			in:   0,
			want: "0",
		},
	}
	for _, tt := range tests {
		got := convertToBase7(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
