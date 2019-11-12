package problem485

func findMaxConsecutiveOnes(nums []int) int {
	ans, count := 0, 0
	for _, v := range nums {
		if v == 1 {
			count++
			if count > ans {
				ans = count
			}
		} else {
			count = 0
		}
	}
	return ans
}
