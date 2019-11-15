package problem1047

import "testing"

type testType struct {
	in   string
	want string
}

func TestRemoveDuplicates(t *testing.T) {
	tests := [...]testType{
		{
			in:   "abbaca",
			want: "ca",
		},
		{
			in:   "aaabaca",
			want: "abaca",
		},
	}
	for _, tt := range tests {
		got := removeDuplicates(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
