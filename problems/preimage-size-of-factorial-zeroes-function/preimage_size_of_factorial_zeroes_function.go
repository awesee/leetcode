package problem793

func preimageSizeFZF(K int) int {
	l, r := 0, 5*(K+1)
	for l <= r {
		m := l + (r-l)/2
		km := zeros(m)
		if km < K {
			l = m + 1
		} else if km > K {
			r = m - 1
		} else {
			return 5
		}
	}
	return 0
}

func zeros(x int) int {
	r := 0
	for x > 0 {
		r += x / 5
		x /= 5
	}
	return r
}
