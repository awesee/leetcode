package problem746

func minCostClimbingStairs(cost []int) int {
	f1, f2, l := 0, 0, len(cost)-1
	for i := l; i >= 0; i-- {
		f0 := f1
		if f2 < f1 {
			f0 = f2
		}
		f0 += cost[i]
		f1, f2 = f0, f1
	}
	if f1 <= f2 {
		return f1
	}
	return f2
}
