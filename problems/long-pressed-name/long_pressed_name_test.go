package problem925

import "testing"

type testType struct {
	name  string
	typed string
	want  bool
}

func TestIsLongPressedName(t *testing.T) {
	tests := [...]testType{
		{
			name:  "alex",
			typed: "aaleex",
			want:  true,
		},
		{
			name:  "saeed",
			typed: "ssaaedd",
			want:  false,
		},
		{
			name:  "leelee",
			typed: "lleeelee",
			want:  true,
		},
		{
			name:  "laiden",
			typed: "laiden",
			want:  true,
		},
		{
			name:  "laiden",
			typed: "laidef",
			want:  false,
		},
		{
			name:  "a",
			typed: "a",
			want:  true,
		},
		{
			name:  "a",
			typed: "b",
			want:  false,
		},
		{
			name:  "pyplrz",
			typed: "ppyypllr",
			want:  false,
		},
		{
			name:  "vtkgn",
			typed: "vttkgnn",
			want:  true,
		},
	}
	for _, tt := range tests {
		got := isLongPressedName(tt.name, tt.typed)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.name, tt.typed, got, tt.want)
		}
	}
}
