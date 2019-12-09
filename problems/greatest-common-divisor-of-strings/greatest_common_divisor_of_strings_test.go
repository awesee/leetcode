package problem1071

import "testing"

type testType struct {
	str1 string
	str2 string
	want string
}

func TestGcdOfStrings(t *testing.T) {
	tests := [...]testType{
		{
			str1: "ABCABC",
			str2: "ABC",
			want: "ABC",
		},
		{
			str1: "ABABAB",
			str2: "ABAB",
			want: "AB",
		},
		{
			str1: "LEET",
			str2: "CODE",
			want: "",
		},
	}
	for _, tt := range tests {
		got := gcdOfStrings(tt.str1, tt.str2)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.str1, got, tt.want)
		}
	}
}
