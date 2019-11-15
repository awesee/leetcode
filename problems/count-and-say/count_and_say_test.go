package problem38

import "testing"

func TestCountAndSay(t *testing.T) {
	tests := map[int]string{
		1: "1",
		3: "21",
		5: "111221",
		7: "13112221",
		9: "31131211131221",
	}

	for in, want := range tests {
		got := countAndSay(in)
		if got != want {
			t.Fatalf("in: %v, got: %v, want: %v", in, got, want)
		}
	}
}
