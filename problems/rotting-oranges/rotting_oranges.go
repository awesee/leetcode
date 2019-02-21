package rotting_oranges

func orangesRotting(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	hasFresh, isRotten, rottens := false, false, make([]int, 0)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if !hasFresh && grid[i][j] == 1 {
				hasFresh = true
			} else if grid[i][j] == 2 {
				rottens = append(rottens, i*c+j)
			}
		}
	}
	for _, p := range rottens {
		i, j := p/c, p%c
		for k := 0; k < 4; k++ {
			nr, nc := i+k-1, j+k-2
			if k == 0 {
				nc += 2
			} else if k == 3 {
				nr -= 2
			}
			if 0 <= nr && nr < r && 0 <= nc && nc < c && grid[nr][nc] == 1 {
				grid[nr][nc], isRotten = 2, true
			}
		}
	}
	if !hasFresh {
		return 0
	} else if !isRotten {
		return -1
	}
	ans := orangesRotting(grid)
	if ans == -1 {
		return -1
	}
	return ans + 1
}
