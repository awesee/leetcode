package problem819

import "testing"

type testType struct {
	paragraph string
	banned    []string
	want      string
}

func TestMostCommonWord(t *testing.T) {
	tests := [...]testType{
		{
			paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:    []string{"hit"},
			want:      "ball",
		},
		{
			paragraph: "a, a, a, a, b,b,b,c, c",
			banned:    []string{"a"},
			want:      "b",
		},
		{
			paragraph: "Bob",
			banned:    []string{},
			want:      "bob",
		},
	}
	for _, tt := range tests {
		got := mostCommonWord(tt.paragraph, tt.banned)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.paragraph, tt.banned, got, tt.want)
		}
	}
}
