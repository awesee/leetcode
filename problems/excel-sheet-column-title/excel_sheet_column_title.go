package problem168

// Solution 1: recursive
func convertToTitle(n int) string {
	if n--; n < 26 {
		return string(n + 'A')
	}
	return convertToTitle(n/26) + string(n%26+'A')
}

// Solution 2: iteration
func convertToTitle2(n int) string {
	ans := ""
	for ; n > 0; n /= 26 {
		n--
		ans = string(n%26+'A') + ans
	}
	return ans
}
