package problem50

func myPow(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			res *= x
		}
		x *= x
	}
	return res
}
