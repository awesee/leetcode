package rotated_digits

func rotatedDigits(N int) int {
	count, mask25, mask69 := 0, 2|5, 6|9
flag:
	for i := 2; i <= N; i++ {
		x, r, base := i, 0, 1
		for x != 0 {
			m := x % 10
			switch m {
			case 2, 5:
				m ^= mask25
			case 6, 9:
				m ^= mask69
			case 3, 4, 7:
				continue flag
			}
			r += m * base
			base *= 10
			x /= 10
		}
		if i != r {
			count++
		}
	}
	return count
}
