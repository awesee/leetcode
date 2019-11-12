package problem383

func canConstruct(ransomNote string, magazine string) bool {
	var m [128]int
	for _, c := range magazine {
		m[c]++
	}
	for _, c := range ransomNote {
		m[c]--
		if m[c] < 0 {
			return false
		}
	}
	return true
}
