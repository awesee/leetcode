package problem942

func diStringMatch(S string) []int {
	lo, hi := 0, len(S)
	ans := make([]int, hi, hi+1)
	for i, v := range S {
		if v == 'I' {
			ans[i] = lo
			lo++
		} else {
			ans[i] = hi
			hi--
		}
	}
	return append(ans, lo)
}
