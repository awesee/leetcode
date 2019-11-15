package problem1023

import (
	"reflect"
	"testing"
)

type testType struct {
	queries []string
	pattern string
	want    []bool
}

func TestCamelMatch(t *testing.T) {
	tests := [...]testType{
		{
			queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern: "FB",
			want:    []bool{true, false, true, true, false},
		},
		{
			queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern: "FoBa",
			want:    []bool{true, false, true, false, false},
		},
		{
			queries: []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"},
			pattern: "FoBaT",
			want:    []bool{false, true, false, false, false},
		},
		{
			queries: []string{"CompetitiveProgramming", "CounterPick", "ControlPanel"},
			pattern: "CooP",
			want:    []bool{false, false, true},
		},
		{
			queries: []string{"aksvbjLiknuTzqon", "ksvjLimflkpnTzqn", "mmkasvjLiknTxzqn", "ksvjLiurknTzzqbn", "ksvsjLctikgnTzqn", "knzsvzjLiknTszqn"},
			pattern: "ksvjLiknTzqn",
			want:    []bool{true, true, true, true, true, true},
		},
	}
	for _, tt := range tests {
		got := camelMatch(tt.queries, tt.pattern)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.pattern, got, tt.want)
		}
	}
}
