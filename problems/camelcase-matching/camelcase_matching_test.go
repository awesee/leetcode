package camelcase_matching

import (
	"reflect"
	"testing"
)

type caseType struct {
	queries  []string
	pattern  string
	expected []bool
}

func TestCamelMatch(t *testing.T) {
	tests := [...]caseType{
		{
			queries:  []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern:  "FB",
			expected: []bool{true, false, true, true, false},
		},
		{
			queries:  []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern:  "FoBa",
			expected: []bool{true, false, true, false, false},
		},
		{
			queries:  []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern:  "FoBaT",
			expected: []bool{false, true, false, false, false},
		},
		{
			queries:  []string{"CompetitiveProgramming", "CounterPick", "ControlPanel"},
			pattern:  "CooP",
			expected: []bool{false, false, true},
		},
		{
			queries:  []string{"aksvbjLiknuTzqon", "ksvjLimflkpnTzqn", "mmkasvjLiknTxzqn", "ksvjLiurknTzzqbn", "ksvsjLctikgnTzqn", "knzsvzjLiknTszqn"},
			pattern:  "ksvjLiknTzqn",
			expected: []bool{true, true, true, true, true, true},
		},
	}
	for _, tc := range tests {
		output := camelMatch(tc.queries, tc.pattern)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.pattern, output, tc.expected)
		}
	}
}
