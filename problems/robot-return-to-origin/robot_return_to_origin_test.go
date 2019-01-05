package robot_return_to_origin

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestJudgeCircle(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "UD",
			expected: true,
		},
		{
			input:    "RL",
			expected: true,
		},
		{
			input:    "LL",
			expected: false,
		},
		{
			input:    "UU",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := judgeCircle(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}

func BenchmarkJudgeCircle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		judgeCircle("RLUDRLUDRLUDRLUDRLUDRLUDRLUDRLUD")
	}
}
