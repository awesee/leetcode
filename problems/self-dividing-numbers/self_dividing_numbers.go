package problem728

func selfDividingNumbers(left int, right int) []int {
	ans := make([]int, 0)
flag:
	for i := left; i <= right; i++ {
		for x := i; x > 0; x /= 10 {
			if d := x % 10; d == 0 || i%d != 0 {
				continue flag
			}
		}
		ans = append(ans, i)
	}
	return ans
}
