package kit

import "testing"

func TestIsEqualSliceInt(t *testing.T) {
	tests := [...]struct {
		s1       []int
		s2       []int
		expected bool
	}{
		{
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 3},
			expected: true,
		},
		{
			s1:       []int{1, 2, 3},
			s2:       []int{3, 2, 1},
			expected: true,
		},
		{
			s1:       []int{1, 2, 3},
			s2:       []int{1, 2, 1},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := IsEqualSliceInt(tc.s1, tc.s2)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.s1, tc.s2, output, tc.expected)
		}
	}
}
