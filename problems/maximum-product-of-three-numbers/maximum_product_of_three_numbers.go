package problem628

import "math"

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt32, math.MaxInt32
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	for _, n := range nums {
		if n < min1 {
			min1, min2 = n, min1
		} else if n < min2 {
			min2 = n
		}
		if n > max3 {
			max1, max2, max3 = max2, max3, n
		} else if n > max2 {
			max1, max2 = max2, n
		} else if n > max1 {
			max1 = n
		}
	}
	ans1, ans2 := min1*min2*max3, max1*max2*max3
	if ans1 > ans2 {
		return ans1
	}
	return ans2
}
