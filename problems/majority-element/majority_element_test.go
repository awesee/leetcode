package problem169

import "testing"

type testType struct {
	in   []int
	want int
}

func TestMajorityElement(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 2, 3},
			want: 3,
		},
		{
			in:   []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			in:   []int{2, 2, 2, 2, 1, 1, 1},
			want: 2,
		},
		{
			in:   []int{1},
			want: 1,
		},
	}
	for _, tt := range tests {
		got := majorityElement(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
