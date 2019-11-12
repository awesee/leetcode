package problem717

func isOneBitCharacter(bits []int) bool {
	i, l := 0, len(bits)-1
	for i < l {
		i += bits[i] + 1
	}
	return i == l
}
