package problem1041

func isRobotBounded(instructions string) bool {
	x, y, dx, dy := 0, 0, 0, 1
	for _, i := range instructions {
		switch i {
		case 'G':
			x, y = x+dx, y+dy
		case 'L':
			dx, dy = -dy, dx
		case 'R':
			dx, dy = dy, -dx
		}
	}
	return x == 0 && y == 0 || dx != 0 || dy != 1
}
