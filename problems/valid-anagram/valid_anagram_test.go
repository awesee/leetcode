package problem242

import "testing"

type testType struct {
	s    string
	t    string
	want bool
}

func TestIsAnagram(t *testing.T) {
	tests := [...]testType{
		{
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			s:    "hannah",
			t:    "hanna",
			want: false,
		},
		{
			s:    "this is string",
			t:    "this is string long",
			want: false,
		},
		{
			s:    "this is string",
			t:    "this as string",
			want: false,
		},
		{
			s:    "你好，世界",
			t:    "世界，你好",
			want: true,
		},
	}

	for _, tt := range tests {
		got := isAnagram(tt.s, tt.t)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.s, tt.t, got, tt.want)
		}
	}
}
