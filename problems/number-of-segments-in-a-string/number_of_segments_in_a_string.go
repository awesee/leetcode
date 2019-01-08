package number_of_segments_in_a_string

import "regexp"

func countSegments(s string) int {
	reg := regexp.MustCompile(`\S+`)
	matches := reg.FindAllString(s, -1)
	return len(matches)
}
