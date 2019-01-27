package k_diff_pairs_in_an_array

func findPairs(nums []int, k int) int {
	ans, m := 0, make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for n, v := range m {
		if k < 0 {
			return ans
		} else if k == 0 && v >= 2 {
			ans++
		} else if k > 0 && m[n+k] > 0 {
			ans++
		}
	}
	return ans
}
