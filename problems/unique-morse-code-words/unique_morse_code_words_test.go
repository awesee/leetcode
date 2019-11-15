package problem804

import "testing"

type testType struct {
	in   []string
	want int
}

func TestUniqueMorseRepresentations(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"gin", "zen", "gig", "msg"},
			want: 2,
		},
		{
			in:   []string{"hello", "word"},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := uniqueMorseRepresentations(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
