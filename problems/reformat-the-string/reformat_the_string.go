package problem1417

func reformat(s string) string {
	i, j, l := 0, 1, len(s)
	str := make([]byte, l+1)
	for _, c := range s {
		if '0' <= c && c <= '9' {
			if i > l {
				return ""
			}
			str[i], i = byte(c), i+2
		} else if j > l {
			return ""
		} else {
			str[j], j = byte(c), j+2
		}
	}
	if i == l || j == l {
		return string(str[:l])
	}
	if i == l-1 {
		str[i] = str[0]
		return string(str[1 : l+1])
	}
	return ""
}
