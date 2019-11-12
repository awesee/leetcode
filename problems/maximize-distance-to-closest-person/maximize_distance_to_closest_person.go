package problem849

func maxDistToClosest(seats []int) int {
	ans, ec := 0, 0
	for i, v := range seats {
		if v == 0 {
			ec++
		} else if ec > 0 {
			if i == ec { // 001
				ans = ec
			} else if (ec+1)/2 > ans { // 101
				ans = (ec + 1) / 2
			}
			ec = 0
		}
	}
	if ec > ans { // 100
		ans = ec
	}
	return ans
}
