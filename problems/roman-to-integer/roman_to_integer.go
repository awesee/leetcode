package problem13

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans, p := 0, 0
	for i := 0; i < len(s); i++ {
		k := s[i]
		if v, ok := m[k]; ok {
			if v > p {
				ans -= p * 2
			}
			ans += v
			p = v
		}
	}
	return ans
}
