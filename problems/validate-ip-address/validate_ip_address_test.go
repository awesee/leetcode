package problem468

import "testing"

type testType struct {
	in   string
	want string
}

func TestValidIPAddress(t *testing.T) {
	tests := [...]testType{
		{
			in:   "172.16.254.1",
			want: "IPv4",
		},
		{
			in:   "172.16.254.01",
			want: "Neither",
		},
		{
			in:   "256.256.256.256",
			want: "Neither",
		},
		{
			in:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			want: "IPv6",
		},
		{
			in:   "2001:db8:85a3:0:0:8A2E:0370:7334",
			want: "IPv6",
		},
		{
			in:   "2001:0db8:85a3::8A2E:0370:7334",
			want: "Neither",
		},
		{
			in:   "02001:0db8:85a3:0000:0000:8a2e:0370:7334",
			want: "Neither",
		},
	}
	for _, tt := range tests {
		got := validIPAddress(tt.in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
