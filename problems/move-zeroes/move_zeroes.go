package problem283

func moveZeroes(nums []int) {
	insertAt := 0
	for i, v := range nums {
		if v != 0 {
			if i != insertAt {
				nums[insertAt] = v
				nums[i] = 0
			}
			insertAt++
		}
	}
}
