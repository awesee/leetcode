package flower_planting_with_no_adjacent

func gardenNoAdj(N int, paths [][]int) []int {
	ans, adjacent, flag := make([]int, N), make([][]int, N), make([][4]bool, N)
	for _, path := range paths {
		if path[0] > path[1] {
			path[0], path[1] = path[1], path[0]
		}
		adjacent[path[0]-1] = append(adjacent[path[0]-1], path[1]-1)
	}
	for i := 0; i < N; i++ {
		for flower, able := range flag[i] {
			if !able {
				ans[i] = flower + 1
				break
			}
		}
		for _, garden := range adjacent[i] {
			flag[garden][ans[i]-1] = true
		}
	}
	return ans
}
