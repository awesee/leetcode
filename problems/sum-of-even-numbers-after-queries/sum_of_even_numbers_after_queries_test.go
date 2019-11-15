package problem985

import (
	"reflect"
	"testing"
)

type testType struct {
	in      []int
	queries [][]int
	want    []int
}

func TestSumEvenAfterQueries(t *testing.T) {
	tests := [...]testType{
		{
			in:      []int{1, 2, 3, 4},
			queries: [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}},
			want:    []int{8, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		got := sumEvenAfterQueries(tt.in, tt.queries)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
