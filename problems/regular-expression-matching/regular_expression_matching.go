package problem10

import "regexp"

func isMatch(s string, p string) bool {
	reg := regexp.MustCompile(p)
	res := reg.FindAllString(s, 1)
	return len(res) == 1 && res[0] == s
}
