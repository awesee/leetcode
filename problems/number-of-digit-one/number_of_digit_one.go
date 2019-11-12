package problem233

func countDigitOne(n int) int {
	ans, b, x := 0, 1, 0
	for b <= n {
		x = n / b
		ans += (x + 8) / 10 * b
		if x%10 == 1 {
			ans += n%b + 1
		}
		b *= 10
	}
	return ans
}
