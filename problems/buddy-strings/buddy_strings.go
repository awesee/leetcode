package problem859

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		exist := [26]bool{}
		for _, c := range A {
			k := c - 'a'
			if exist[k] {
				return true
			}
			exist[k] = true
		}
		return false
	}
	m, n := -1, -1
	for i, c := range A {
		if B[i] != byte(c) {
			if m == -1 {
				m = i
			} else if n == -1 {
				n = i
			} else {
				return false
			}
		}
	}
	return n != -1 && A[m] == B[n] && A[n] == B[m]
}
