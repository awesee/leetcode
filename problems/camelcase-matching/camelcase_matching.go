package problem1023

func camelMatch(queries []string, pattern string) []bool {
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ptn := []rune(pattern)
		for _, r := range q {
			if len(ptn) > 0 && ptn[0] == r {
				ptn = ptn[1:]
			} else if r >= 'A' && r <= 'Z' {
				ptn = []rune{1}
				break
			}
		}
		ans[i] = len(ptn) == 0
	}
	return ans
}
