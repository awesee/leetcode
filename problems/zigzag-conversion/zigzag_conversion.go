package problem6

func convert(s string, numRows int) string {
	l := len(s)
	if l <= numRows || numRows < 2 {
		return s
	}
	ans, col := make([]byte, 0, l), (l-1)/(numRows-1)
	for r := 0; r < numRows; r++ {
		for c := 0; c <= col; c++ {
			if c&1 == 0 || r > 0 && r < numRows-1 {
				k := (numRows-1)*(c+c&1) + r*(1-c&1*2)
				if k < l {
					ans = append(ans, s[k])
				}
			}
		}
	}
	return string(ans)
}
