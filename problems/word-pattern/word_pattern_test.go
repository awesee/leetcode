package problem290

import "testing"

type testType struct {
	pattern string
	str     string
	want    bool
}

func TestWordPattern(t *testing.T) {
	tests := [...]testType{
		{
			pattern: "abba",
			str:     "dog cat cat dog",
			want:    true,
		},
		{
			pattern: "abba",
			str:     "dog cat cat fish",
			want:    false,
		},
		{
			pattern: "aaaa",
			str:     "dog cat cat dog",
			want:    false,
		},
		{
			pattern: "abba",
			str:     "dog dog dog dog",
			want:    false,
		},
		{
			pattern: "abba",
			str:     "dog cat cat dog fish",
			want:    false,
		},
		{
			pattern: "abba",
			str:     "b a a b",
			want:    true,
		},
		{
			pattern: "abba",
			str:     "b c c b",
			want:    true,
		},
	}
	for _, tt := range tests {
		got := wordPattern(tt.pattern, tt.str)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.pattern, got, tt.want)
		}
	}
}
