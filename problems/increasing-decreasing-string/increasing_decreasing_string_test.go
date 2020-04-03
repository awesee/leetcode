package problem1370

import "testing"

type testType struct {
	in   string
	want string
}

func TestSortString(t *testing.T) {
	tests := [...]testType{
		{
			in:   "aaaabbbbcccc",
			want: "abccbaabccba",
		},
		{
			in:   "rat",
			want: "art",
		},
		{
			in:   "leetcode",
			want: "cdelotee",
		},
		{
			in:   "ggggggg",
			want: "ggggggg",
		},
		{
			in:   "spo",
			want: "ops",
		},
	}
	for _, tt := range tests {
		got := sortString(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
