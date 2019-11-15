package problem650

import "testing"

type testType struct {
	in   int
	want int
}

func TestMinSteps(t *testing.T) {
	tests := [...]testType{
		{
			in:   1,
			want: 0,
		},
		{
			in:   3,
			want: 3,
		},
		{
			in:   30,
			want: 10,
		},
		{
			in:   97,
			want: 97,
		},
	}
	for _, tt := range tests {
		got := minSteps(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
