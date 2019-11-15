package problem1029

import "testing"

type testType struct {
	in   [][]int
	want int
}

func TestTwoCitySchedCost(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{10, 20},
				{30, 200},
				{400, 50},
				{30, 20},
			},
			want: 110,
		},
		{
			in: [][]int{
				{259, 770},
				{448, 54},
				{926, 667},
				{184, 139},
				{840, 118},
				{577, 469},
			},
			want: 1859,
		},
		{
			in: [][]int{
				{918, 704},
				{77, 778},
				{239, 457},
				{284, 263},
				{872, 779},
				{976, 416},
				{860, 518},
				{13, 351},
				{137, 238},
				{557, 596},
				{890, 227},
				{548, 143},
				{384, 585},
				{165, 54},
			},
			want: 4532,
		},
	}
	for _, tt := range tests {
		got := twoCitySchedCost(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
