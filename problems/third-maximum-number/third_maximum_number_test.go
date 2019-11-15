package problem414

import "testing"

type testType struct {
	in   []int
	want int
}

func TestThirdMax(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 2, 5, 3, 5},
			want: 2,
		},
		{
			in:   []int{1, 2, 2},
			want: 2,
		},
		{
			in:   []int{1, 1, 2},
			want: 2,
		},
		{
			in:   []int{3, 2, 1},
			want: 1,
		},
		{
			in:   []int{1, 2},
			want: 2,
		},
		{
			in:   []int{2, 2, 3, 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		got := thirdMax(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
