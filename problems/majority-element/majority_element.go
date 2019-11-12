package problem169

func majorityElement(nums []int) int {
	ans, count := 0, 0
	for _, n := range nums {
		if n == ans {
			count++
		} else {
			count--
		}
		if count <= 0 {
			ans = n
			count = 1
		}
	}
	return ans
}
