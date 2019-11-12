package problem17

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}
	m := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	ans := make([]string, 1)
	for _, k := range digits {
		idx, t := 0, make([]string, len(ans))
		copy(t, ans)
		for _, s := range t {
			for _, c := range m[k] {
				if idx < len(ans) {
					ans[idx] = s + c
				} else {
					ans = append(ans, s+c)
				}
				idx++
			}
		}
	}
	return ans
}
