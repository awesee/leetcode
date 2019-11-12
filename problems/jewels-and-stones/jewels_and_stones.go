package problem771

func numJewelsInStones(J string, S string) int {
	m := [128]bool{}
	for i := 0; i < len(J); i++ {
		m[J[i]] = true
	}
	count := 0
	for i := 0; i < len(S); i++ {
		if m[S[i]] {
			count++
		}
	}
	return count
}
