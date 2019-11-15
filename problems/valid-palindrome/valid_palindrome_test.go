package problem125

import "testing"

type testType struct {
	in   string
	want bool
}

func TestIsPalindrome(t *testing.T) {
	tests := [...]testType{
		{
			in:   "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			in:   "",
			want: true,
		},
		{
			in:   ".",
			want: true,
		},
		{
			in:   "12321",
			want: true,
		},
		{
			in:   "race a car",
			want: false,
		},
		{
			in:   "hello, world",
			want: false,
		},
	}

	for _, tt := range tests {
		got := isPalindrome(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
