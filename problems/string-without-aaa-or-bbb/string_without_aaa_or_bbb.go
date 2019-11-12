package problem984

func strWithout3a3b(A int, B int) string {
	str, ac, bc := "", 0, 0
	for A > 0 || B > 0 {
		if A > B && ac < 2 || bc == 2 {
			str += "a"
			A, ac, bc = A-1, ac+1, 0
		} else {
			str += "b"
			B, ac, bc = B-1, 0, bc+1
		}
	}
	return str
}
