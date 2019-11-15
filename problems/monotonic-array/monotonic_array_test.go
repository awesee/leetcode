package problem896

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestIsMonotonic(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 2, 3},
			want: true,
		},
		{
			in:   []int{6, 5, 4, 4},
			want: true,
		},
		{
			in:   []int{1, 3, 2},
			want: false,
		},
		{
			in:   []int{1, 2, 4, 5},
			want: true,
		},
		{
			in:   []int{1, 1, 1},
			want: true,
		},
		{
			in:   []int{1, 2, 3, 3, 3, 2, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isMonotonic(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
