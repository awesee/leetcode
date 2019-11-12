package problem1051

import "sort"

func heightChecker(heights []int) int {
	ans, dst := 0, make([]int, len(heights))
	copy(dst, heights)
	sort.Ints(dst)
	for i, v := range dst {
		if heights[i] != v {
			ans++
		}
	}
	return ans
}
