package problem788

func rotatedDigits(N int) int {
	count := 0
	for i := 2; i <= N; i++ {
		x, ok := i, 0
		for x != 0 {
			switch x % 10 {
			case 2, 5, 6, 9:
				ok = 1
			case 3, 4, 7:
				ok, x = 0, 0
			}
			x /= 10
		}
		count += ok
	}
	return count
}
