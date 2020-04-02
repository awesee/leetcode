package problem1374

import "testing"

type testType struct {
	in   int
	want string
}

func TestGenerateTheString(t *testing.T) {
	tests := [...]testType{
		{
			in:   4,
			want: "xxxy",
		},
		{
			in:   2,
			want: "xy",
		},
		{
			in:   7,
			want: "xxxxxxx",
		},
		{
			in:   30,
			want: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxy",
		},
	}
	for _, tt := range tests {
		got := generateTheString(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
