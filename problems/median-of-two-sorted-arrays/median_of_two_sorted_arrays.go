package problem4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	s := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		v1 := nums1[i]
		v2 := nums2[j]
		if v1 <= v2 {
			s = append(s, v1)
			i++
		} else {
			s = append(s, v2)
			j++
		}
	}
	if i < m {
		s = append(s, nums1[i:]...)
	} else {
		s = append(s, nums2[j:]...)
	}
	count := m + n
	if count&1 == 1 {
		k := (count - 1) / 2
		return float64(s[k])
	}
	k := count / 2
	return (float64(s[k-1]) + float64(s[k])) / 2
}
