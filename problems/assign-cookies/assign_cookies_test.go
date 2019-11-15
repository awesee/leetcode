package problem455

import "testing"

type testType struct {
	g    []int
	s    []int
	want int
}

func TestFindContentChildren(t *testing.T) {
	tests := [...]testType{
		{
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			want: 1,
		},
		{
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := findContentChildren(tt.g, tt.s)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.g, got, tt.want)
		}
	}
}
