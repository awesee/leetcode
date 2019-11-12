package problem994

func orangesRotting(grid [][]int) int {
	r, c, minutes := len(grid), len(grid[0]), 0
	freshCount, rottens := 0, make([]int, 0)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				freshCount++
			} else if grid[i][j] == 2 {
				rottens = append(rottens, i*c+j)
			}
		}
	}
	for l := len(rottens); l > 0 && freshCount > 0; l = len(rottens) {
		for _, p := range rottens {
			for k := 0; k < 4; k++ {
				i, j := p/c+k-2, p%c+k-1
				if k == 0 {
					i += 2
				} else if k == 3 {
					j -= 2
				}
				if 0 <= i && i < r && 0 <= j && j < c && grid[i][j] == 1 {
					grid[i][j], freshCount, rottens = 2, freshCount-1, append(rottens, i*c+j)
				}
			}
		}
		minutes, rottens = minutes+1, rottens[l:]
	}
	if freshCount > 0 {
		return -1
	}
	return minutes
}
