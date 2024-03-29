package problem929

import "testing"

type testType struct {
	in   []string
	want int
}

func TestNumUniqueEmails(t *testing.T) {
	tests := [...]testType{
		{
			in:   []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.ttode.com"},
			want: 2,
		},
		{
			in:   []string{"awesee+wang@awesee.cn", "awesee.wang@awesee.cn", "aweSee+wang@awesee.cn"},
			want: 3,
		},
	}

	for _, tt := range tests {
		got := numUniqueEmails(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
