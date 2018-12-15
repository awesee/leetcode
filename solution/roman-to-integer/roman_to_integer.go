package roman_to_integer

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	r := 0
	p := 0
	l := len(s)
	for i := 0; i < l; i++ {
		k := s[i:i+1]
		if v, ok := m[k]; ok {
			if v > p {
				r -= p * 2
			}
			r += v
			p = v
		}
	}

	return r
}
