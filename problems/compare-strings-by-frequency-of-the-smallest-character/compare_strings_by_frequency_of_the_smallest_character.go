package problem1170

import "sort"

func numSmallerByFrequency(queries []string, words []string) []int {
	ans, m := make([]int, len(queries)), make([]int, len(words))
	for i, w := range words {
		m[i] = f(w)
	}
	sort.Ints(m)
	for i, q := range queries {
		t := f(q)
		for j := len(m) - 1; j >= 0; j-- {
			if t >= m[j] {
				break
			}
			ans[i]++
		}
	}
	return ans
}

func f(s string) int {
	m, p := [26]int{}, byte(25)
	for i := 0; i < len(s); i++ {
		k := s[i] - 'a'
		m[k]++
		if p > k {
			p = k
		}
	}
	return m[p]
}
