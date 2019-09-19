package problem_1185

import "testing"

type caseType struct {
	day      int
	month    int
	year     int
	expected string
}

func TestDayOfTheWeek(t *testing.T) {
	tests := [...]caseType{
		{
			day:      31,
			month:    8,
			year:     2019,
			expected: "Saturday",
		},
		{
			day:      15,
			month:    8,
			year:     1993,
			expected: "Sunday",
		},
		{
			day:      1,
			month:    1,
			year:     1970,
			expected: "Thursday",
		},
		{
			day:      25,
			month:    3,
			year:     1993,
			expected: "Thursday",
		},
		{
			day:      30,
			month:    6,
			year:     2000,
			expected: "Friday",
		},
	}
	for _, tc := range tests {
		output := dayOfTheWeek(tc.day, tc.month, tc.year)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc, output, tc.expected)
		}
	}
}
