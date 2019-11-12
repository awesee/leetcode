package problem443

import "strconv"

func compress(chars []byte) int {
	ans, anchor, l := 0, 0, len(chars)
	for i, c := range chars {
		if i == l-1 || c != chars[i+1] {
			chars[ans] = chars[anchor]
			ans++
			if i > anchor {
				for _, n := range strconv.Itoa(i - anchor + 1) {
					chars[ans] = byte(n)
					ans++
				}
			}
			anchor = i + 1
		}
	}
	return ans
}
