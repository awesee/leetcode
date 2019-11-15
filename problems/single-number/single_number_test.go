package problem136

import "testing"

type testType struct {
	in   []int
	want int
}

func TestSingleNumber(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{2, 2, 1},
			want: 1,
		},
		{
			in:   []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			in:   []int{1, 2, 1, 2, 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		got := singleNumber(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
