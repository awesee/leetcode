package problem997

import "testing"

type testType struct {
	N     int
	trust [][]int
	want  int
}

func TestFindJudge(t *testing.T) {
	tests := [...]testType{
		{
			N: 2,
			trust: [][]int{
				{1, 2},
			},
			want: 2,
		},
		{
			N: 3,
			trust: [][]int{
				{1, 3},
				{2, 3},
			},
			want: 3,
		},
		{
			N: 3,
			trust: [][]int{
				{1, 3},
				{2, 3},
				{3, 1},
			},
			want: -1,
		},
		{
			N: 3,
			trust: [][]int{
				{1, 2},
				{2, 3},
			},
			want: -1,
		},
		{
			N: 4,
			trust: [][]int{
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{4, 3},
			},
			want: 3,
		},
		{
			N:     1,
			trust: [][]int{},
			want:  1,
		},
	}
	for _, tt := range tests {
		got := findJudge(tt.N, tt.trust)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.N, tt.trust, got, tt.want)
		}
	}
}
