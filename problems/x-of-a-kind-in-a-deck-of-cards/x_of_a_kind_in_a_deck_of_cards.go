package x_of_a_kind_in_a_deck_of_cards

func hasGroupsSizeX(deck []int) bool {
	m, l := make(map[int]int), len(deck)
	for _, v := range deck {
		m[v]++
	}
flag:
	for x := 2; x <= l; x++ {
		if l%x == 0 {
			for _, v := range m {
				if v%x != 0 {
					continue flag
				}
			}
			return true
		}
	}
	return false
}
