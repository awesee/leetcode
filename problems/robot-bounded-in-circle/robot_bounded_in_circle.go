package robot_bounded_in_circle

func isRobotBounded(instructions string) bool {
	x, y, p := 0, 0, 0
	for _, i := range instructions {
		switch i {
		case 'G':
			if p == 1 || p == 3 {
				x += 2 - p
			} else {
				y += 1 - p
			}
		case 'L':
			p = (p + 3) % 4
		case 'R':
			p = (p + 1) % 4
		}
	}
	return p != 0 || x == 0 && y == 0
}
