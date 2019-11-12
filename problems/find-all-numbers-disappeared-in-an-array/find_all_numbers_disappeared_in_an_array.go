package problem448

func findDisappearedNumbers(nums []int) []int {
	max, count := len(nums), 0
	s := make([]bool, max)
	for _, v := range nums {
		s[v-1] = true
	}
	for i := 1; i <= max; i++ {
		if !s[i-1] {
			nums[count] = i
			count++
		}
	}
	return nums[:count]
}
