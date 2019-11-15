package problem67

import "testing"

type testType struct {
	a    string
	b    string
	want string
}

func TestAddBinary(t *testing.T) {
	tests := [...]testType{
		{
			a:    "11",
			b:    "1",
			want: "100",
		},
		{
			a:    "1010",
			b:    "1011",
			want: "10101",
		},
		{
			a:    "111",
			b:    "111",
			want: "1110",
		},
	}
	for _, tt := range tests {
		got := addBinary(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
