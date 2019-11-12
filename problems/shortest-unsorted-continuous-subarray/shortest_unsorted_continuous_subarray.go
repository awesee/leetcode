package problem581

func findUnsortedSubarray(nums []int) int {
	m, n, l := 0, -1, len(nums)
	for max, min, i, j := 0, l-1, 1, l-2; i < l; i, j = i+1, j-1 {
		if nums[i] < nums[max] {
			n = i
		} else if nums[i] > nums[max] {
			max = i
		}
		if nums[j] > nums[min] {
			m = j
		} else if nums[j] < nums[min] {
			min = j
		}
	}
	return n - m + 1
}
