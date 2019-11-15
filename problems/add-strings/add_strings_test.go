package problem415

import "testing"

type testType struct {
	num1 string
	num2 string
	want string
}

func TestAddStrings(t *testing.T) {
	tests := [...]testType{
		{
			num1: "0",
			num2: "0",
			want: "0",
		},
		{
			num1: "1",
			num2: "2",
			want: "3",
		},
		{
			num1: "9",
			num2: "9",
			want: "18",
		},
		{
			num1: "100",
			num2: "999",
			want: "1099",
		},
	}
	for _, tt := range tests {
		got := addStrings(tt.num1, tt.num2)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt, got, tt.want)
		}
	}
}
