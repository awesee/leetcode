package problem16

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans, diff, n := math.MaxInt32, math.MaxInt32, len(nums)
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				if target-sum < diff {
					ans, diff = sum, target-sum
				}
				l++
			} else if sum > target {
				if sum-target < diff {
					ans, diff = sum, sum-target
				}
				r--
			} else {
				return target
			}
		}
	}
	return ans
}
