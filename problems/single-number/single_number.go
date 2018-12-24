package single_number

func singleNumber(nums []int) int {
	n := 0
	for _, v := range nums {
		n ^= v
	}
	return n
}
