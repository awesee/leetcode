package problem1436

import "testing"

type testType struct {
	in   [][]string
	want string
}

func TestDestCity(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]string{
				{"London", "New York"},
				{"New York", "Lima"},
				{"Lima", "Sao Paulo"},
			},
			want: "Sao Paulo",
		},
		{
			in: [][]string{
				{"B", "C"},
				{"D", "B"},
				{"C", "A"},
			},
			want: "A",
		},
		{
			in: [][]string{
				{"A", "Z"},
			},
			want: "Z",
		},
	}
	for _, tt := range tests {
		got := destCity(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
