package problem1374

func generateTheString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = 'x'
	}
	if n%2 == 0 {
		b[n-1] = 'y'
	}
	return string(b)
}
