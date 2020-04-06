package problem1399

import "testing"

type testType struct {
	in   int
	want int
}

func TestCountLargestGroup(t *testing.T) {
	tests := [...]testType{
		{
			in:   13,
			want: 4,
		},
		{
			in:   2,
			want: 2,
		},
		{
			in:   15,
			want: 6,
		},
		{
			in:   24,
			want: 5,
		},
	}
	for _, tt := range tests {
		got := countLargestGroup(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
