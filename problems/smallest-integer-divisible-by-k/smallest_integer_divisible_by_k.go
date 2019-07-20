package smallest_integer_divisible_by_k

func smallestRepunitDivByK(K int) int {
	if K%2 == 0 || K%5 == 0 {
		return -1
	}

	r, length := 1%K, 1
	for r != 0 && length <= K {
		r = (r*10 + 1) % K
		length++
	}

	return length
}
