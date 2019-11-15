package problem944

import "testing"

type testType struct {
	in   []string
	want int
}

func TestMinDeletionSize(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"cba", "daf", "ghi"},
			want: 1,
		},
		{
			in:   []string{"a", "b"},
			want: 0,
		},
		{
			in:   []string{"zyx", "wvu", "tsr"},
			want: 3,
		},
	}
	for _, tt := range tests {
		got := minDeletionSize(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
