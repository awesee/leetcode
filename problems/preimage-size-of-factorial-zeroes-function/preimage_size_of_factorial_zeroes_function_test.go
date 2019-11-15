package problem793

import "testing"

type testType struct {
	in   int
	want int
}

func TestPreimageSizeFZF(t *testing.T) {
	tests := [...]testType{
		{
			in:   0,
			want: 5,
		},
		{
			in:   5,
			want: 0,
		},
		{
			in:   17,
			want: 0,
		},
		{
			in:   11,
			want: 0,
		},
	}
	for _, tt := range tests {
		got := preimageSizeFZF(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
