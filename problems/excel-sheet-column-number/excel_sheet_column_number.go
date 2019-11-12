package problem171

func titleToNumber(s string) int {
	ans := 0
	for _, c := range s {
		ans *= 26
		ans += int(c) - 'A' + 1
	}
	return ans
}
