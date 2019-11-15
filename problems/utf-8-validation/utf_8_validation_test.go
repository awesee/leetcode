package problem393

import "testing"

type testType struct {
	in   []int
	want bool
}

func TestValidUtf8(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{197},
			want: false,
		},
		{
			in:   []int{197, 130, 1},
			want: true,
		},
		{
			in:   []int{235, 140, 4},
			want: false,
		},
		{
			in:   []int{248},
			want: false,
		},
		{
			in:   []int{194, 155, 231, 184, 185, 246, 176, 131, 161, 222, 174, 227, 162, 134, 241, 154, 168, 185, 218, 178, 229, 187, 139, 246, 178, 187, 139, 204, 146, 225, 148, 179, 245, 139, 172, 134, 193, 156, 233, 131, 154, 240, 166, 188, 190, 216, 150, 230, 145, 144, 240, 167, 140, 163, 221, 190, 238, 168, 139, 241, 154, 159, 164, 199, 170, 224, 173, 140, 244, 182, 143, 134, 206, 181, 227, 172, 141, 241, 146, 159, 170, 202, 134, 230, 142, 163, 244, 172, 140, 191},
			want: true,
		},
	}
	for _, tt := range tests {
		got := validUtf8(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
