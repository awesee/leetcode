package problem412

import (
	"reflect"
	"testing"
)

type testType struct {
	in   int
	want []string
}

func TestFizzBuzz(t *testing.T) {
	tests := [...]testType{
		{
			in: 15,
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz",
			},
		},
		{
			in: 5,
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
			},
		},
		{
			in: 10,
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
			},
		},
	}

	for _, tt := range tests {
		got := fizzBuzz(tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
