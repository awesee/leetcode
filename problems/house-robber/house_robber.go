package problem198

func rob(nums []int) int {
	pre1, pre2, ans := 0, 0, 0
	for _, v := range nums {
		if ans = pre1 + v; pre2 > ans {
			ans = pre2
		}
		pre1, pre2 = pre2, ans
	}
	return ans
}
