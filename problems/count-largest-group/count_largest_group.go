package problem1399

func countLargestGroup(n int) int {
	group := make(map[int]int)
	for i := 1; i <= n; i++ {
		s := 0
		for x := i; x > 0; x /= 10 {
			s += x % 10
		}
		group[s]++
	}
	ans, max := 0, 0
	for _, v := range group {
		if v == max {
			ans++
		} else if v > max {
			ans, max = 1, v
		}
	}
	return ans
}
