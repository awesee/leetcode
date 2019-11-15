package problem1108

import "testing"

type testType struct {
	in   string
	want string
}

func TestDefangIPaddr(t *testing.T) {
	tests := [...]testType{
		{
			in:   "1.1.1.1",
			want: "1[.]1[.]1[.]1",
		},
		{
			in:   "255.100.50.0",
			want: "255[.]100[.]50[.]0",
		},
	}
	for _, tt := range tests {
		got := defangIPaddr(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
