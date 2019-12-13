package problem507

// Solution 1
func checkPerfectNumber(num int) bool {
	if num < 6 {
		return false
	}
	sum, mid := 1, num
	for i := 2; i < mid; i++ {
		if num%i == 0 {
			n := num / i
			sum += i + n
			if mid > n {
				mid = n
			}
		}
	}
	return sum == num
}

// Solution 2
func checkPerfectNumber2(num int) bool {
	m := map[int]bool{
		6:        true,
		28:       true,
		496:      true,
		8128:     true,
		33550336: true,
	}
	return m[num]
}
