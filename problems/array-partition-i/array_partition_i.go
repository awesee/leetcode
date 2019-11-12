package problem561

import "sort"

func arrayPairSum(nums []int) int {
	sum := 0
	sort.Ints(nums)
	for i, v := range nums {
		if i&1 == 0 {
			sum += v
		}
	}
	return sum
}
