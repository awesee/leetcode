package problem66

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
