package problem1309

import "testing"

type testType struct {
	in   string
	want string
}

func TestFreqAlphabets(t *testing.T) {
	tests := [...]testType{
		{
			in:   "10#11#12",
			want: "jkab",
		},
		{
			in:   "1326#",
			want: "acz",
		},
		{
			in:   "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#",
			want: "abcdefghijklmnopqrstuvwxyz",
		},
	}
	for _, tt := range tests {
		got := freqAlphabets(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
