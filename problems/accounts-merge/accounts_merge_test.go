package problem721

import (
	"reflect"
	"sort"
	"testing"
)

type testType struct {
	in   [][]string
	want [][]string
}

func TestAccountsMerge(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			want: [][]string{
				{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			in: [][]string{
				{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
				{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
				{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
				{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
			},
			want: [][]string{
				{"Alex", "Alex0@m.co", "Alex4@m.co", "Alex5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan3@m.co"},
				{"Gabe", "Gabe0@m.co", "Gabe2@m.co", "Gabe3@m.co", "Gabe4@m.co"},
				{"Kevin", "Kevin2@m.co", "Kevin4@m.co"},
			},
		},
		{
			in: [][]string{
				{"David", "David0@m.co", "David4@m.co", "David3@m.co"},
				{"David", "David5@m.co", "David5@m.co", "David0@m.co"},
				{"David", "David1@m.co", "David4@m.co", "David0@m.co"},
				{"David", "David0@m.co", "David1@m.co", "David3@m.co"},
				{"David", "David4@m.co", "David1@m.co", "David3@m.co"},
			},
			want: [][]string{
				{"David", "David0@m.co", "David1@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
			},
		},
		{
			in: [][]string{
				{"David", "David0@m.co", "David1@m.co"},
				{"David", "David3@m.co", "David4@m.co"},
				{"David", "David4@m.co", "David5@m.co"},
				{"David", "David2@m.co", "David3@m.co"},
				{"David", "David1@m.co", "David2@m.co"},
			},
			want: [][]string{
				{"David", "David0@m.co", "David1@m.co", "David2@m.co", "David3@m.co", "David4@m.co", "David5@m.co"},
			},
		},
	}
	for _, tt := range tests {
		got := accountsMerge(tt.in)
		sort.Slice(got, func(i, j int) bool {
			if got[i][0] == got[i][0] {
				return got[i][1] < got[j][1]
			}
			return got[i][0] < got[i][0]
		})
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
