package problem383

import "testing"

type testType struct {
	ransom   string
	magazine string
	want     bool
}

func TestCanConstruct(t *testing.T) {
	tests := [...]testType{
		{
			ransom:   "a",
			magazine: "b",
			want:     false,
		},
		{
			ransom:   "aa",
			magazine: "ab",
			want:     false,
		},
		{
			ransom:   "aa",
			magazine: "aab",
			want:     true,
		},
	}
	for _, tt := range tests {
		got := canConstruct(tt.ransom, tt.magazine)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.ransom, tt.magazine, got, tt.want)
		}
	}
}
