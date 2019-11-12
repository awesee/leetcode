package problem7

import "math"

func reverse(x int) int {
	r := 0
	for x != 0 {
		r = r*10 + x%10
		if r > math.MaxInt32 || r < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return r
}
