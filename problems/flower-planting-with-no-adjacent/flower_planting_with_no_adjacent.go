package problem1042

func gardenNoAdj(N int, paths [][]int) []int {
	ans, adjGarden, flowerUsed := make([]int, N), make([][]int, N), make([][4]bool, N)
	for _, path := range paths {
		if path[0] > path[1] {
			path[0], path[1] = path[1], path[0]
		}
		adjGarden[path[0]-1] = append(adjGarden[path[0]-1], path[1]-1)
	}
	for i := 0; i < N; i++ {
		for flower, used := range flowerUsed[i] {
			if !used {
				ans[i] = flower + 1
				break
			}
		}
		for _, garden := range adjGarden[i] {
			flowerUsed[garden][ans[i]-1] = true
		}
	}
	return ans
}
