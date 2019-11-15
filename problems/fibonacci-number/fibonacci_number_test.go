package problem509

import "testing"

type testType struct {
	in   int
	want int
}

func TestFib(t *testing.T) {
	tests := [...]testType{
		{
			in:   2,
			want: 1,
		},
		{
			in:   3,
			want: 2,
		},
		{
			in:   4,
			want: 3,
		},
		{
			in:   1,
			want: 1,
		},
	}
	for _, tt := range tests {
		got := fib(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
