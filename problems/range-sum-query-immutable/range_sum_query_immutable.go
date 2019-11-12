package problem303

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	sum, na := 0, NumArray{make([]int, len(nums)+1)}
	na.sum[0] = 0
	for i, v := range nums {
		sum += v
		na.sum[i+1] = sum
	}
	return na
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sum[j+1] - this.sum[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
