package problem674

func findLengthOfLCIS(nums []int) int {
	max, cur := 0, 1
	for i, v := range nums {
		if i >= 1 && v > nums[i-1] {
			cur++
		} else {
			cur = 1
		}
		if cur > max {
			max = cur
		}
	}
	return max
}
