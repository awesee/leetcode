package problem859

import "testing"

type testType struct {
	a    string
	b    string
	want bool
}

func TestBuddyStrings(t *testing.T) {
	tests := [...]testType{
		{
			a:    "ab",
			b:    "ba",
			want: true,
		},
		{
			a:    "aa",
			b:    "aa",
			want: true,
		},
		{
			a:    "ab",
			b:    "ab",
			want: false,
		},
		{
			a:    "aaaaaaabc",
			b:    "aaaaaaacb",
			want: true,
		},
		{
			a:    "",
			b:    "aa",
			want: false,
		},
		{
			a:    "hello",
			b:    "h0lle",
			want: false,
		},
		{
			a:    "hello",
			b:    "hanna",
			want: false,
		},
	}
	for _, tt := range tests {
		got := buddyStrings(tt.a, tt.b)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.a, tt.b, got, tt.want)
		}
	}
}
