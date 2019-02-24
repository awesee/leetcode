package power_of_three

func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}
