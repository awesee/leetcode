package problem69

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	r := x / 2
	for r > x/r {
		r = (r + x/r) / 2
	}
	return r
}
