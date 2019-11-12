package problem888

func fairCandySwap(A []int, B []int) []int {
	sA, sB := 0, 0
	mB := make(map[int]bool)
	for _, v := range A {
		sA += v
	}
	for _, v := range B {
		sB += v
		mB[v] = true
	}
	for _, x := range A {
		if mB[x+(sB-sA)/2] {
			return []int{x, x + (sB-sA)/2}
		}
	}
	return nil
}
