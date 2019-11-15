package problem877

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestStoneGame(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{5, 3, 4, 5},
			want: true,
		},
		{
			in:   []int{2, 5, 7, 3},
			want: true,
		},
	}
	for _, tt := range tests {
		got := stoneGame(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
