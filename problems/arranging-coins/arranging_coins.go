package problem441

import "math"

func arrangeCoins(n int) int {
	x := math.Sqrt(1+8*float64(n)) - 1
	return int(x) / 2
}
