package problem724

func pivotIndex(nums []int) int {
	left, right := 0, 0
	for _, v := range nums {
		right += v
	}
	for i, v := range nums {
		if right -= v; left == right {
			return i
		}
		left += v
	}
	return -1
}
