package problem71

import "testing"

type testType struct {
	in   string
	want string
}

func TestSimplifyPath(t *testing.T) {
	tests := [...]testType{
		{
			in:   "/home/",
			want: "/home",
		},
		{
			in:   "/../",
			want: "/",
		},
		{
			in:   "/home//foo/",
			want: "/home/foo",
		},
		{
			in:   "/a/./b/../../c/",
			want: "/c",
		},
		{
			in:   "/a/../../b/../c//.//",
			want: "/c",
		},
		{
			in:   "/a//b////c/d//././/..",
			want: "/a/b/c",
		},
	}
	for _, tt := range tests {
		got := simplifyPath(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
