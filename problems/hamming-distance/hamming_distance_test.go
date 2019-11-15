package problem461

import "testing"

type testType struct {
	x    int
	y    int
	want int
}

func TestHammingDistance(t *testing.T) {
	tests := [...]testType{
		{
			x:    1,
			y:    4,
			want: 2,
		},
	}
	for _, tt := range tests {
		got := hammingDistance(tt.x, tt.y)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.x, tt.y, got, tt.want)
		}
	}
}
