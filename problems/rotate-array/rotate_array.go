package problem189

func rotate(nums []int, k int) {
	l := len(nums)
	k %= l
	s := make([]int, l)
	copy(s, nums)
	for i, v := range s[l-k:] {
		nums[i] = v
	}
	for i, v := range s[:l-k] {
		nums[k+i] = v
	}
}
