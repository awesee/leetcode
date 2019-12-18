package problem343

func integerBreak(n int) int {
	if n <= 3 {
		return n - 1
	}
	ans := 1
	for n > 4 {
		n -= 3
		ans *= 3
	}
	return ans * n
}
