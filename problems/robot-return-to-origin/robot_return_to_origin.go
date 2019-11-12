package problem657

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, move := range moves {
		switch move {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
	}
	return x == 0 && y == 0
}
