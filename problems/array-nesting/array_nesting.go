package problem565

func arrayNesting(nums []int) int {
	ans := 0
	for i := range nums {
		c := 1
		for nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			c++
		}
		if ans < c {
			ans = c
		}
	}
	return ans
}
