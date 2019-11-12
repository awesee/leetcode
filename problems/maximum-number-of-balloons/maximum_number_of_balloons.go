package problem1189

func maxNumberOfBalloons(text string) int {
	ans, m := len(text), make(map[byte]int, 5)
	for i := 0; i < ans; i++ {
		switch text[i] {
		case 'b', 'a', 'l', 'o', 'n':
			m[text[i]]++
		}
	}
	for _, c := range [...]byte{'b', 'a', 'l', 'o', 'n'} {
		if c == 'l' || c == 'o' {
			m[c] /= 2
		}
		if ans > m[c] {
			ans = m[c]
		}
	}
	return ans
}
