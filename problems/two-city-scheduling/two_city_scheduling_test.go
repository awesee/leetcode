package two_city_scheduling

import "testing"

type caseType struct {
	input    [][]int
	expected int
}

func TestTwoCitySchedCost(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{10, 20},
				{30, 200},
				{400, 50},
				{30, 20},
			},
			expected: 110,
		},
		{
			input: [][]int{
				{259, 770},
				{448, 54},
				{926, 667},
				{184, 139},
				{840, 118},
				{577, 469},
			},
			expected: 1859,
		},
		{
			input: [][]int{
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
			expected: 4532,
		},
	}
	for _, tc := range tests {
		output := twoCitySchedCost(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
