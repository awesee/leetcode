package valid_palindrome_ii

func validPalindrome(s string) bool {
	start, end := 0, len(s)-1
	i, j, times := 0, 0, 0
	for start < end {
		if s[start] == s[end] {
			start++
			end--
		} else {
			if times == 0 {
				times++
				i, j = start, end
				start++
			} else if times == 1 {
				times++
				start, end = i, j-1
			} else {
				return false
			}
		}
	}
	return true
}
