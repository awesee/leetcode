package problem268

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMissingNumber(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 0, 1},
			want: 2,
		},
		{
			in:   []int{0, 1, 2, 4, 5, 6},
			want: 3,
		},
		{
			in:   []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			want: 8,
		},
	}
	for _, tt := range tests {
		got := missingNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
