package problem665

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestCheckPossibility(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 2, 5},
			want: true,
		},
		{
			in:   []int{3, 2, 1},
			want: false,
		},
		{
			in:   []int{3, 4, 2, 3},
			want: false,
		},
		{
			in:   []int{1, 2, 3, 7, 5, 6},
			want: true,
		},
		{
			in:   []int{1, 2, 3, 7, 1, 9},
			want: true,
		},
	}
	for _, tt := range tests {
		got := checkPossibility(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
