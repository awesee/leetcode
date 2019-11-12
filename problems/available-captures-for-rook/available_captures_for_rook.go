package problem999

func numRookCaptures(board [][]byte) int {
	ans, rx, ry, left, up := 0, -1, -1, [8]int{}, [8]int{}
	for i, row := range board {
		for j, v := range row {
			if rx == -1 {
				if v == 'p' {
					left[i], up[j] = 1, 1
				} else if v == 'B' {
					left[i], up[j] = 0, 0
				} else if v == 'R' {
					ans, rx, ry = left[i]+up[j], i, j
				}
			} else if i == rx {
				if v == 'p' {
					ans++
				}
				if v != '.' {
					break
				}
			} else if j == ry {
				if v == 'p' {
					ans++
				}
				if v != '.' {
					return ans
				}
			}
		}
	}
	return ans
}
