package problem509

func fib(N int) int {
	a, b := 0, 1
	for i := 0; i < N; i++ {
		a, b = b, a+b
	}
	return a
}
