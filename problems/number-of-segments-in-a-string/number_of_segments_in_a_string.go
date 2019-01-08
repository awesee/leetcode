package number_of_segments_in_a_string

import "strings"

func countSegments(s string) int {
	return len(strings.Fields(s))
}
