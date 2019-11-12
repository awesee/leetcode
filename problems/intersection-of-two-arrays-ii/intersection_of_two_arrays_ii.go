package problem350

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1))
	nc := make(map[int]int)
	for _, n := range nums2 {
		nc[n]++
	}
	for _, n := range nums1 {
		if nc[n] > 0 {
			res = append(res, n)
			nc[n]--
		}
	}
	return res
}
