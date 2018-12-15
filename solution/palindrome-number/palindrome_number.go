package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	rx, t := 0, 0
	n := x
	for x != 0 {
		t = rx*10 + x%10
		if t/10 != rx {
			return false
		}
		rx = t
		x /= 10
	}
	return rx == n
}
