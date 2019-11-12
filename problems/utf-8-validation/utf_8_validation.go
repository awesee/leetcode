package problem393

func validUtf8(data []int) bool {
	u := 0
	for _, v := range data {
		if u == 0 && v>>7 == 0 { // 0b0
			u = 0
		} else if u == 0 && v>>5 == 6 { // 0b110
			u = 1
		} else if u == 0 && v>>4 == 14 { // 0b1110
			u = 2
		} else if u == 0 && v>>3 == 30 { // 0b11110
			u = 3
		} else if u > 0 && v>>6 == 2 { // 0b10
			u--
		} else {
			return false
		}
	}
	return u == 0
}
