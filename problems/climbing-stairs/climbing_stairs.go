package problem70

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first, second := 1, 2
	for i := 3; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}
