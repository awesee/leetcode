package problem893

func numSpecialEquivGroups(A []string) int {
	m := make(map[[52]int]bool)
	for _, s := range A {
		var count [52]int
		for i, char := range s {
			count[int(char)-'a'+26*(i%2)]++
		}
		m[count] = true
	}
	return len(m)
}
