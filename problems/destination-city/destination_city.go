package problem1436

func destCity(paths [][]string) string {
	m := make(map[string]bool)
	for _, path := range paths {
		m[path[0]] = true
	}
	for _, path := range paths {
		if !m[path[1]] {
			return path[1]
		}
	}
	return ""
}
