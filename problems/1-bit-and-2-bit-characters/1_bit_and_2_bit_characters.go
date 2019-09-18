package problem_717

func isOneBitCharacter(bits []int) bool {
	l := len(bits) - 1
	i := 0
	for i < l {
		i += bits[i] + 1
	}
	return i == l
}
