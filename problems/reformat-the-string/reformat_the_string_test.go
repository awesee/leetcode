package problem1417

import "testing"

type testType struct {
	in   string
	want string
}

func TestReformat(t *testing.T) {
	tests := [...]testType{
		{
			in:   "a0b1c2",
			want: "0a1b2c",
		},
		{
			in:   "leetcode",
			want: "",
		},
		{
			in:   "1229857369",
			want: "",
		},
		{
			in:   "covid2019",
			want: "c0o1v9i2d",
		},
		{
			in:   "ab123",
			want: "1a2b3",
		},
		{
			in:   "ec",
			want: "",
		},
		{
			in:   "abcd12345",
			want: "1a2b3c4d5",
		},
		{
			in:   "12345abcd",
			want: "1a2b3c4d5",
		},
		{
			in:   "77",
			want: "",
		},
		{
			in:   "1",
			want: "1",
		},
		{
			in:   "a",
			want: "a",
		},
	}
	for _, tt := range tests {
		got := reformat(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
