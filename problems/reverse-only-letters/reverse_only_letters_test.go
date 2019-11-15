package problem917

import "testing"

type testType struct {
	in   string
	want string
}

func TestReverseOnlyLetters(t *testing.T) {
	tests := [...]testType{
		{
			in:   "ab-cd",
			want: "dc-ba",
		},
		{
			in:   "a-bC-dEf-ghIj",
			want: "j-Ih-gfE-dCba",
		},
		{
			in:   "Test1ng-Leet=code-Q!",
			want: "Qedo1ct-eeLg=ntse-T!",
		},
	}
	for _, tt := range tests {
		got := reverseOnlyLetters(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
