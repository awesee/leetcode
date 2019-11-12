package problem461

func hammingDistance(x int, y int) int {
	ans, n := 0, x^y
	for n > 0 {
		ans, n = ans+1, n&(n-1)
	}
	return ans
}
