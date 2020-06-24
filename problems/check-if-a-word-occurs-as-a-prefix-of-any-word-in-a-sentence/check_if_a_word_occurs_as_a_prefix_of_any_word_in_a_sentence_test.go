package problem1455

import "testing"

type testType struct {
	sentence   string
	searchWord string
	want       int
}

func TestIsPrefixOfWord(t *testing.T) {
	tests := [...]testType{
		{
			sentence:   "i love eating burger",
			searchWord: "burg",
			want:       4,
		},
		{
			sentence:   "this problem is an easy problem",
			searchWord: "pro",
			want:       2,
		},
		{
			sentence:   "i am tired",
			searchWord: "you",
			want:       -1,
		},
		{
			sentence:   "i use triple pillow",
			searchWord: "pill",
			want:       4,
		},
		{
			sentence:   "hello from the other side",
			searchWord: "they",
			want:       -1,
		},
	}
	for _, tt := range tests {
		got := isPrefixOfWord(tt.sentence, tt.searchWord)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.sentence, got, tt.want)
		}
	}
}
