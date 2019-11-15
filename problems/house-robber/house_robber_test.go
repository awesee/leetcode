package problem198

import "testing"

type testType struct {
	in   []int
	want int
}

func TestRob(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 1},
			want: 4,
		},
		{
			in:   []int{2, 7, 9, 3, 1},
			want: 12,
		},
	}
	for _, tt := range tests {
		got := rob(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
