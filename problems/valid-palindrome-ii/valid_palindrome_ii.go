package problem680

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	i, j, times := 0, 0, 0
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			times++
			if times == 1 {
				i, j = start, end
				start++
			} else if times == 2 {
				start, end = i, j-1
			} else {
				return false
			}
		}
	}
	return true
}
