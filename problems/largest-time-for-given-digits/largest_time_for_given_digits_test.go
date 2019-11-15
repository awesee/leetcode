package problem949

import "testing"

type testType struct {
	in   []int
	want string
}

func TestLargestTimeFromDigits(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 4},
			want: "23:41",
		},
		{
			in:   []int{5, 5, 5, 5},
			want: "",
		},
		{
			in:   []int{2, 0, 6, 6},
			want: "06:26",
		},
	}
	for _, tt := range tests {
		got := largestTimeFromDigits(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
