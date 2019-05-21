package camelcase_matching

import "regexp"

func camelMatch(queries []string, pattern string) []bool {
	ans, ptn := make([]bool, len(queries)), []rune("[a-z]*")
	for _, r := range pattern {
		ptn = append(ptn, r)
		ptn = append(ptn, []rune("[a-z]*")...)
	}
	reg := regexp.MustCompile(string(ptn))
	for i, q := range queries {
		ans[i] = reg.FindString(q) == q
	}
	return ans
}
