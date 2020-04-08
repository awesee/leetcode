package problem1389

func createTargetArray(nums []int, index []int) []int {
	ans := make([]int, len(nums))
	for i, idx := range index {
		copy(ans[idx+1:], ans[idx:])
		ans[idx] = nums[i]
	}
	return ans
}
