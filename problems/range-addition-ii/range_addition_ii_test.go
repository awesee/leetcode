package problem598

import "testing"

type testType struct {
	m    int
	n    int
	ops  [][]int
	want int
}

func TestMaxCount(t *testing.T) {
	tests := [...]testType{
		{
			m: 3,
			n: 3,
			ops: [][]int{
				{2, 2},
				{3, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		got := maxCount(tt.m, tt.n, tt.ops)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.m, got, tt.want)
		}
	}
}
