package problem665

func checkPossibility(nums []int) bool {
	modified := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if modified {
				return false
			} else if i >= 1 && nums[i-1] > nums[i+1] {
				nums[i+1] = nums[i]
			}
			modified = true
		}
	}
	return true
}
