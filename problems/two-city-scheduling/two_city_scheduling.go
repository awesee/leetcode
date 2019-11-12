package problem1029

import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.Slice(costs, func(i, j int) bool {
		return abs(costs[i][0]-costs[i][1]) > abs(costs[j][0]-costs[j][1])
	})
	cost, a, b, n := 0, 0, 0, len(costs)/2
	for _, c := range costs {
		if (c[0] < c[1] && a < n) || b == n {
			cost += c[0]
			a++
		} else {
			cost += c[1]
			b++
		}
	}
	return cost
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
