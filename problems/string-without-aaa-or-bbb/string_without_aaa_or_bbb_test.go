package problem984

import "testing"

type testType struct {
	a    int
	b    int
	want string
}

func TestStrWithout3a3b(t *testing.T) {
	tests := [...]testType{
		{
			a:    6,
			b:    2,
			want: "aabaabaa",
		},
		{
			a:    2,
			b:    6,
			want: "bbabbabb",
		},
		{
			a:    1,
			b:    2,
			want: "bba",
		},
		{
			a:    4,
			b:    1,
			want: "aabaa",
		},
	}
	for _, tt := range tests {
		got := strWithout3a3b(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
