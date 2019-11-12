package problem414

func thirdMax(nums []int) int {
	max1, max2, max3 := 0, -1, -1
	for i, v := range nums {
		if v > nums[max1] {
			max1, max2, max3 = i, max1, max2
		} else if (max2 == -1 || v > nums[max2]) && v != nums[max1] {
			max2, max3 = i, max2
		} else if (max3 == -1 || v > nums[max3]) && v != nums[max1] && v != nums[max2] {
			max3 = i
		}
	}
	if max3 == -1 {
		return nums[max1]
	}
	return nums[max3]
}
