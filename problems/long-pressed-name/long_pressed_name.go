package problem925

func isLongPressedName(name string, typed string) bool {
	i, nl, j, tl := 0, len(name), 0, len(typed)
	for i < nl || j < tl {
		if i < nl && j < tl && name[i] == typed[j] {
			i++
			j++
		} else if i >= 1 && j < tl && typed[j] == name[i-1] {
			j++
		} else {
			return false
		}
	}
	return true
}
