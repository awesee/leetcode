package problem122

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMaxProfit(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{7, 1, 5, 3, 6, 4},
			want: 7,
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			in:   []int{7, 6, 4, 3, 1},
			want: 0,
		},
	}

	for _, tt := range tests {
		got := maxProfit(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
