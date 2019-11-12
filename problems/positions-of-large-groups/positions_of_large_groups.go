package problem830

func largeGroupPositions(S string) [][]int {
	pre, ans, l := 0, make([][]int, 0), len(S)
	for i, v := range S {
		if i == l-1 || v != rune(S[i+1]) {
			if i-pre >= 2 {
				ans = append(ans, []int{pre, i})
			}
			pre = i + 1
		}
	}
	return ans
}
