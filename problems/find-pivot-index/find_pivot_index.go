package find_pivot_index

func pivotIndex(nums []int) int {
	left, right := 0, 0
	for _, v := range nums {
		right += v
	}
	for i, v := range nums {
		right -= v
		if left == right {
			return i
		} else {
			left += v
		}
	}
	return -1
}
