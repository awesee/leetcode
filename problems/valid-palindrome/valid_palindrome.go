package problem125

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if !isAlphanumeric(s[start]) {
			start++
			continue
		}
		if !isAlphanumeric(s[end]) {
			end--
			continue
		}
		x := s[start]
		if s[start] >= 'a' {
			x = s[start] + 'A' - 'a'
		}
		y := s[end]
		if s[end] >= 'a' {
			y = s[end] + 'A' - 'a'
		}
		if x != y {
			return false
		}
		start++
		end--
	}
	return true
}

func isAlphanumeric(b byte) bool {
	return ('0' <= b && b <= '9') || ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z')
}
