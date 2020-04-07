package problem1394

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	ans, max := -1, 0
	for k, v := range m {
		if k > max && k == v {
			ans, max = k, v
		}
	}
	return ans
}
