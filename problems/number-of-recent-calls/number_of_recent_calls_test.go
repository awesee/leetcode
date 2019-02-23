package number_of_recent_calls

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestConstructor(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 100, 3001, 3002},
			expected: []int{1, 2, 3, 3},
		},
	}
	for _, tc := range tests {
		obj, output := Constructor(), make([]int, 0)
		for _, t := range tc.input {
			output = append(output, obj.Ping(t))
		}
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
