package problem747

func dominantIndex(nums []int) int {
	ans, sec := 0, 0
	for i, v := range nums {
		if v > nums[ans] {
			ans, sec = i, nums[ans]
		} else if v > sec && i != ans {
			sec = v
		}
	}
	if nums[ans] >= 2*sec {
		return ans
	}
	return -1
}
