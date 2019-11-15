package problem191

import "testing"

type testType struct {
	in   uint32
	want int
}

func TestHammingWeight(t *testing.T) {
	tests := [...]testType{
		{
			in:   3,
			want: 2,
		},
		{
			in:   7,
			want: 3,
		},
		{
			in:   8,
			want: 1,
		},
		{
			in:   0x5555,
			want: 8,
		},
	}
	for _, tt := range tests {
		got := hammingWeight(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}

func BenchmarkHammingWeight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hammingWeight(0x55555555)
	}
}
