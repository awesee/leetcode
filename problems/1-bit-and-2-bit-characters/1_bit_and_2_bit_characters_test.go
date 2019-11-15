package problem717

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestIsOneBitCharacter(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 0, 0},
			want: true,
		},
		{
			in:   []int{1, 1, 1, 0},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isOneBitCharacter(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
