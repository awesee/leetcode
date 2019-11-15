package problem824

import "testing"

type testType struct {
	in   string
	want string
}

func TestToGoatLatin(t *testing.T) {
	tests := [...]testType{
		{
			in:   "I speak Goat Latin",
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			in:   "The quick brown fox jumped over the lazy dog",
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		got := toGoatLatin(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
