package problem88

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[n:], nums1)
	cur, i, j := 0, n, 0
	for j < n {
		if i < m+n && nums1[i] < nums2[j] {
			nums1[cur], i = nums1[i], i+1
		} else {
			nums1[cur], j = nums2[j], j+1
		}
		cur++
	}
}
