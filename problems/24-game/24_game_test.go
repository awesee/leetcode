package problem679

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestJudgePoint24(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{4, 1, 8, 7},
			want: true,
		},
		{
			in:   []int{1, 2, 1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		got := judgePoint24(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
