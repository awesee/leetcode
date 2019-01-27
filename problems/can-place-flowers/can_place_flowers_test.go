package can_place_flowers

import "testing"

type caseType struct {
	flowerbed []int
	n         int
	expected  bool
}

func TestCanPlaceFlowers(t *testing.T) {
	tests := [...]caseType{
		{
			flowerbed: []int{0, 1, 0, 0, 0, 0, 1, 0},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{1, 0, 0, 0, 1},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{0, 0, 1, 0},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{0, 1, 0, 0},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{0, 1, 0, 0, 1, 0},
			n:         1,
			expected:  false,
		},
		{
			flowerbed: []int{0, 0, 0, 0, 0},
			n:         3,
			expected:  true,
		},
		{
			flowerbed: []int{0, 0},
			n:         2,
			expected:  false,
		},
		{
			flowerbed: []int{0},
			n:         1,
			expected:  true,
		},
		{
			flowerbed: []int{0, 1},
			n:         1,
			expected:  false,
		},
	}
	for _, tc := range tests {
		output := canPlaceFlowers(tc.flowerbed, tc.n)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.flowerbed, tc.n, output, tc.expected)
		}
	}
}
