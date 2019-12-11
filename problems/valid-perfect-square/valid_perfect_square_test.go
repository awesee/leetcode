package problem367

import "testing"

type testType struct {
	in   int
	want bool
}

func TestIsPerfectSquare(t *testing.T) {
	tests := [...]testType{
		{
			in:   16,
			want: true,
		},
		{
			in:   14,
			want: false,
		},
		{
			in:   0,
			want: true,
		},
		{
			in:   1,
			want: true,
		},
		{
			in:   3,
			want: false,
		},
	}
	for _, tt := range tests {
		got := isPerfectSquare(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
