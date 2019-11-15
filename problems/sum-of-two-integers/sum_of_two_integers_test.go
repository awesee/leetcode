package problem371

import "testing"

type testType struct {
	a    int
	b    int
	want int
}

func TestGetSum(t *testing.T) {
	tests := [...]testType{
		{
			a:    1,
			b:    2,
			want: 3,
		},
		{
			a:    -2,
			b:    3,
			want: 1,
		},
	}
	for _, tt := range tests {
		got := getSum(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
