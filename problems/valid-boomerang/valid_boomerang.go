package problem1037

func isBoomerang(points [][]int) bool {
	ans := points[0][0] * (points[1][1] - points[2][1])
	ans += points[1][0] * (points[2][1] - points[0][1])
	ans += points[2][0] * (points[0][1] - points[1][1])
	return ans != 0
}
