package problem165

import "testing"

type testType struct {
	v1   string
	v2   string
	want int
}

func TestCompareVersion(t *testing.T) {
	tests := [...]testType{
		{
			v1:   "1.0.01",
			v2:   "1.00.1",
			want: 0,
		},
		{
			v1:   "1.0.0",
			v2:   "1",
			want: 0,
		},
		{
			v1:   "1.01",
			v2:   "1.001",
			want: 0,
		},
		{
			v1:   "1.0",
			v2:   "1.0.0",
			want: 0,
		},
		{
			v1:   "0.1",
			v2:   "1.1",
			want: -1,
		},
		{
			v1:   "7.5.2.4",
			v2:   "7.5.3",
			want: -1,
		},

		{
			v1:   "7.00",
			v2:   "7.001",
			want: -1,
		},
		{
			v1:   "1.1",
			v2:   "1.10",
			want: -1,
		},
		{
			v1:   "1.0.1",
			v2:   "1",
			want: 1,
		},
		{
			v1:   "1.10",
			v2:   "1.2",
			want: 1,
		},
		{
			v1:   "1.10.05",
			v2:   "1.9.10",
			want: 1,
		},
	}
	for _, tt := range tests {
		got := compareVersion(tt.v1, tt.v2)
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.v1, tt.v2, got, tt.want)
		}
	}
}
