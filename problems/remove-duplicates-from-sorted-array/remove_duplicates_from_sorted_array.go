package problem26

func removeDuplicates(nums []int) int {
	top, l := 0, len(nums)
	if l == 0 {
		return 0
	}
	for i := 1; i < l; i++ {
		if nums[i] != nums[top] {
			top++
			nums[top] = nums[i]
		}
	}
	return top + 1
}
