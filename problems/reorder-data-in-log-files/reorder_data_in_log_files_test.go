package problem937

import (
	"reflect"
	"testing"
)

type testType struct {
	in   []string
	want []string
}

func TestReorderLogFiles(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			want: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			in:   []string{"8 fj dzz k", "5r 446 6 3", "63 gu psub", "5 ba iedrz", "6s 87979 5", "3r 587 01", "jc 3480612", "bb wsrd kp", "b aq cojj", "r5 6316 71"},
			want: []string{"b aq cojj", "5 ba iedrz", "8 fj dzz k", "63 gu psub", "bb wsrd kp", "5r 446 6 3", "6s 87979 5", "3r 587 01", "jc 3480612", "r5 6316 71"},
		},
		{
			in:   []string{"qi ir oo i", "cp vnzw i", "0 fkbikbts", "4 j trouka", "gn j q al", "5r w wgqc", "m8 x haje", "fg 28694 6", "i gf mwdoa", "ao 0850716"},
			want: []string{"0 fkbikbts", "i gf mwdoa", "qi ir oo i", "gn j q al", "4 j trouka", "cp vnzw i", "5r w wgqc", "m8 x haje", "fg 28694 6", "ao 0850716"},
		},
	}
	for _, tt := range tests {
		got := make([]string, len(tt.in))
		copy(got, tt.in)
		got = reorderLogFiles(got)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
