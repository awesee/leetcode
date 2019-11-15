package problem605

import "testing"

type testType struct {
	flowerbed []int
	n         int
	want      bool
}

func TestCanPlaceFlowers(t *testing.T) {
	tests := [...]testType{
		{
			flowerbed: []int{0, 1, 0, 0, 0, 0, 1, 0},
			n:         2,
			want:      false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			want:      true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			want:      false,
		},
		{
			flowerbed: []int{0, 0, 1, 0},
			n:         1,
			want:      true,
		},
		{
			flowerbed: []int{0, 1, 0, 0},
			n:         1,
			want:      true,
		},
		{
			flowerbed: []int{0, 1, 0, 0, 1, 0},
			n:         1,
			want:      false,
		},
		{
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         3,
			want:      true,
		},
		{
			flowerbed: []int{0, 0},
			n:         2,
			want:      false,
		},
		{
			flowerbed: []int{0},
			n:         1,
			want:      true,
		},
		{
			flowerbed: []int{0, 1},
			n:         1,
			want:      false,
		},
	}
	for _, tt := range tests {
		got := canPlaceFlowers(tt.flowerbed, tt.n)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.flowerbed, tt.n, got, tt.want)
		}
	}
}
