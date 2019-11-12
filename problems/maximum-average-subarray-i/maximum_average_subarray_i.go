package problem643

func findMaxAverage(nums []int, k int) float64 {
	max, sum := 0, 0
	for i, v := range nums {
		sum += v
		if i == k-1 {
			max = sum
		} else if i >= k {
			sum -= nums[i-k]
		}
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
