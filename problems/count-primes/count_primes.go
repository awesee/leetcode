package problem204

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	primes := make([]bool, n)
	count := 1
	for i := 3; i < n; i += 2 {
		if !primes[i] {
			count++
			for j := i * i; j < n; j += i + i {
				primes[j] = true
			}
		}
	}
	return count
}
