package problem18

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans, end := make([][]int, 0), len(nums)-1
	for i := 0; i < end-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < end-1; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l, r := j+1, end
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum < target {
					l++
				} else if sum > target {
					r--
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[l], nums[r]})
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
