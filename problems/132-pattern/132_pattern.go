package problem456

func find132pattern(nums []int) bool {
	ak, ajStack := -1<<31, make([]int, 0, len(nums))
	for i := len(nums) - 1; 0 <= i; i-- {
		if nums[i] < ak {
			return true
		}
		for len(ajStack) > 0 && ajStack[len(ajStack)-1] < nums[i] {
			ak = ajStack[len(ajStack)-1]
			ajStack = ajStack[:len(ajStack)-1]
		}
		ajStack = append(ajStack, nums[i])
	}
	return false
}
