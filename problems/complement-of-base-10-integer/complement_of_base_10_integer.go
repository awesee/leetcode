package problem1009

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	ans, base := 0, 1
	for N > 0 {
		if N&1 == 0 {
			ans += base
		}
		base, N = base<<1, N>>1
	}
	return ans
}
