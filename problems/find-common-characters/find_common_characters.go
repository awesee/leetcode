package problem1002

func commonChars(A []string) []string {
	ans, m, l := make([]string, 0), make([][26]int, len(A)), len(A)
	for i, str := range A {
		for _, c := range str {
			m[i][c-'a']++
		}
	}
	for i, c := range m[0] {
		for c > 0 {
			c--
			for j := 1; j < l; j++ {
				if m[j][i] > 0 {
					m[j][i]--
				} else {
					c = -1
					break
				}
			}
			if c >= 0 {
				ans = append(ans, string(i+'a'))
			}
		}
	}
	return ans
}
