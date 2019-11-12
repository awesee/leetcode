package problem53

func maxSubArray(nums []int) int {
	max, current, l := nums[0], nums[0], len(nums)
	for i := 1; i < l; i++ {
		if current > 0 {
			current += nums[i]
		} else {
			current = nums[i]
		}
		if current > max {
			max = current
		}
	}
	return max
}
