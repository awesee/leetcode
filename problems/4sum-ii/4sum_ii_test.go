package problem454

import "testing"

type testType struct {
	a    []int
	b    []int
	c    []int
	d    []int
	want int
}

func TestFourSumCount(t *testing.T) {
	tests := [...]testType{
		{
			a:    []int{1, 2},
			b:    []int{-2, -1},
			c:    []int{-1, 2},
			d:    []int{0, 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := fourSumCount(tt.a, tt.b, tt.c, tt.d)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt, got, tt.want)
		}
	}
}
