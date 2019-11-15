package problem657

import "testing"

type testType struct {
	in   string
	want bool
}

func TestJudgeCircle(t *testing.T) {
	tests := [...]testType{
		{
			in:   "UD",
			want: true,
		},
		{
			in:   "RL",
			want: true,
		},
		{
			in:   "LL",
			want: false,
		},
		{
			in:   "UU",
			want: false,
		},
	}
	for _, tt := range tests {
		got := judgeCircle(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}

func BenchmarkJudgeCircle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeCircle("RLUDRLUDRLUDRLUDRLUDRLUDRLUDRLUD")
	}
}
