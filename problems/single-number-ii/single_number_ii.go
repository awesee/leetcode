package problem137

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, n := range nums {
		ones = (ones ^ n) & ^twos
		twos = (twos ^ n) & ^ones
	}
	return ones
}
