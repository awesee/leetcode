package problem15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans, length := make([][]int, 0), len(nums)
	for i := 0; i < length-2; i++ {
		if nums[i] <= 0 && (i == 0 || nums[i] != nums[i-1]) {
			l, r := i+1, length-1
			for l < r {
				sum := nums[i] + nums[l] + nums[r]
				if sum < 0 {
					l++
				} else if sum > 0 {
					r--
				} else {
					ans = append(ans, []int{nums[i], nums[l], nums[r]})
					l, r = l+1, r-1
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}
	return ans
}
