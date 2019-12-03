package problem1170

func numSmallerByFrequency(queries []string, words []string) []int {
	ans, m := make([]int, len(queries)), make([]int, len(words))
	for i, w := range words {
		m[i] = f(w)
	}
	for i, q := range queries {
		t := f(q)
		for _, n := range m {
			if t < n {
				ans[i]++
			}
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
