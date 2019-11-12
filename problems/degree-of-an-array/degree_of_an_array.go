package problem697

func findShortestSubArray(nums []int) int {
	m, ans := make(map[int]int), 0
	first, max := make(map[int]int), 0
	for i, v := range nums {
		m[v]++
		if first[v] == 0 {
			first[v] = i + 1
		}
		if m[v] >= max {
			t := i - first[v] + 2
			if m[v] > max || (m[v] == max && t < ans) {
				ans = t
			}
			max = m[v]
		}
	}
	return ans
}
