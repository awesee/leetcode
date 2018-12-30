package majority_element

func majorityElement(nums []int) int {
	count, ans := 1, nums[0]
	for _, n := range nums[1:] {
		if n != ans {
			count--
		}
		if count <= 0 {
			ans = n
			count = 1
		}
	}
	return ans
}
