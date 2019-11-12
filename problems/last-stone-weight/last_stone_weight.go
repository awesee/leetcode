package problem1046

import "sort"

func lastStoneWeight(stones []int) int {
	for len(stones) >= 2 {
		n := len(stones) - 1
		sort.Ints(stones)
		v := stones[n] - stones[n-1]
		if v > 0 {
			stones[n-1] = v
		} else {
			n--
		}
		stones = stones[:n]
	}
	stones = append(stones, 0)
	return stones[0]
}
