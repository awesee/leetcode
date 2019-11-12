package problem3

func lengthOfLongestSubstring(s string) int {
	ans, l := 0, len(s)
	index := [128]int{}
	for i, j := 0, 0; j < l; j++ {
		if i < index[s[j]] {
			i = index[s[j]]
		}
		if ans < j-i+1 {
			ans = j - i + 1
		}
		index[s[j]] = j + 1
	}
	return ans
}
