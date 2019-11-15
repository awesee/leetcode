package problem303

import "testing"

type testType struct {
	i    int
	j    int
	want int
}

func TestConstructor(t *testing.T) {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := Constructor(nums)
	tests := [...]testType{
		{
			i:    0,
			j:    2,
			want: 1,
		},
		{
			i:    2,
			j:    5,
			want: -1,
		},
		{
			i:    0,
			j:    5,
			want: -3,
		},
	}
	for _, tt := range tests {
		got := obj.SumRange(tt.i, tt.j)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.i, tt.j, got, tt.want)
		}
	}
}
