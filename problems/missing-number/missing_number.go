package problem268

func missingNumber(nums []int) int {
	sum, n := 0, len(nums)
	for _, v := range nums {
		sum += v
	}
	return (n+1)*n/2 - sum
}
