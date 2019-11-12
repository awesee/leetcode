package problem219

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if p, ok := m[v]; ok && i-p <= k {
			return true
		}
		m[v] = i
	}
	return false
}
