package problem28

import "testing"

type testType struct {
	haystack string
	needle   string
	want     int
}

func TestStrStr(t *testing.T) {
	tests := [...]testType{
		{
			haystack: "hello",
			needle:   "ll",
			want:     2,
		}, {
			haystack: "this is test string",
			needle:   "a",
			want:     -1,
		},
	}

	for _, tt := range tests {
		got := strStr(tt.haystack, tt.needle)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.haystack, tt.needle, got, tt.want)
		}
	}
}
