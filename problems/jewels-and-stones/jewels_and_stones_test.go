package problem771

import "testing"

type testType struct {
	j    string
	s    string
	want int
}

func TestNumJewelsInStones(t *testing.T) {
	tests := [...]testType{
		{
			j:    "aA",
			s:    "aAAbbbb",
			want: 3,
		},
		{
			j:    "z",
			s:    "ZZ",
			want: 0,
		},
		{
			j:    "Tb",
			s:    "abcdefhelloaabads",
			want: 2,
		},
	}

	for _, tt := range tests {
		got := numJewelsInStones(tt.j, tt.s)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.j, tt.s, got, tt.want)
		}
	}
}
