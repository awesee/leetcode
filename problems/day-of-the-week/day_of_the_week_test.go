package problem1185

import "testing"

type testType struct {
	day   int
	month int
	year  int
	want  string
}

func TestDayOfTheWeek(t *testing.T) {
	tests := [...]testType{
		{
			day:   31,
			month: 8,
			year:  2019,
			want:  "Saturday",
		},
		{
			day:   15,
			month: 8,
			year:  1993,
			want:  "Sunday",
		},
		{
			day:   1,
			month: 1,
			year:  1970,
			want:  "Thursday",
		},
		{
			day:   25,
			month: 3,
			year:  1993,
			want:  "Thursday",
		},
		{
			day:   30,
			month: 6,
			year:  2000,
			want:  "Friday",
		},
	}
	for _, tt := range tests {
		got := dayOfTheWeek(tt.day, tt.month, tt.year)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt, got, tt.want)
		}
	}
}
