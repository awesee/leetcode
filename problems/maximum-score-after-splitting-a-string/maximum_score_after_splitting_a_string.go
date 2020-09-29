package problem1422

import "strings"

func maxScore(s string) int {
	ans := 0
	n0 := 0
	n1 := strings.Count(s, "1")
	for _, v := range s[:len(s)-1] {
		if v == '1' {
			n1--
		} else {
			n0++
		}
		if ans < n0+n1 {
			ans = n0 + n1
		}
	}
	return ans
}
