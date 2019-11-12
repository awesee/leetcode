package problem202

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true
		sum := 0
		for n > 0 {
			x := n % 10
			sum += x * x
			n /= 10
		}
		n = sum
	}
	return n == 1
}
