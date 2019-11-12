package problem455

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var ans, i, j int
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			ans, i = ans+1, i+1
		}
		j++
	}
	return ans
}
