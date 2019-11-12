package problem744

func nextGreatestLetter(letters []byte, target byte) byte {
	for _, v := range letters {
		if v > target {
			return v
		}
	}
	return letters[0]
}
