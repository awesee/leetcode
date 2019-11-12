package problem970

import "math"

func powerfulIntegers(x int, y int, bound int) []int {
	xFloat, yFloat, boundFloat := float64(x), float64(y), float64(bound)
	ans, m := make([]int, 0), make(map[float64]bool)
	for i := 0.0; i < 18 && math.Pow(xFloat, i) <= boundFloat; i++ {
		for j := 0.0; j < 18 && math.Pow(yFloat, j) <= boundFloat; j++ {
			if v := math.Pow(xFloat, i) + math.Pow(yFloat, j); v <= boundFloat && !m[v] {
				ans, m[v] = append(ans, int(v)), true
			}
		}
	}
	return ans
}
