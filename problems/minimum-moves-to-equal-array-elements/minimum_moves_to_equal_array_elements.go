package problem453

func minMoves(nums []int) int {
	sum, min, n := 0, nums[0], len(nums)
	for _, num := range nums {
		sum += num
		if min > num {
			min = num
		}
	}
	return sum - n*min
}
