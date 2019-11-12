package problem136

func singleNumber(nums []int) int {
	n := 0
	for _, v := range nums {
		n ^= v
	}
	return n
}
