package problem36

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		row := make(map[byte]bool)
		col := make(map[byte]bool)
		boxes := make(map[byte]bool)

		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}
			if board[j][i] != '.' {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}

			if board[i/3*3+j/3][i%3*3+j%3] != '.' {
				if boxes[board[i/3*3+j/3][i%3*3+j%3]] {
					return false
				}
				boxes[board[i/3*3+j/3][i%3*3+j%3]] = true
			}
		}
	}
	return true
}
