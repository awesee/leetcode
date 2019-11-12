package problem1221

func balancedStringSplit(s string) int {
	ans, flag := 0, 0
	for _, b := range s {
		if b == 'L' {
			flag--
		} else {
			flag++
		}
		if flag == 0 {
			ans++
		}
	}
	return ans
}
