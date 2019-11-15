package problem961

import "testing"

type testType struct {
	in   []int
	want int
}

func TestRepeatedNTimes(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 3},
			want: 3,
		},
		{
			in:   []int{2, 1, 2, 5, 3, 2},
			want: 2,
		},
		{
			in:   []int{5, 1, 5, 2, 5, 3, 5, 4},
			want: 5,
		},
		{
			in:   []int{1, 2, 3},
			want: 0,
		},
	}

	for _, tt := range tests {
		got := repeatedNTimes(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
