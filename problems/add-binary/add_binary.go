package add_binary

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	var carry byte = '0'
	var bs []byte
	for i >= 0 || j >= 0 || carry != '0' {
		v := carry
		if i >= 0 {
			v += a[i] - '0'
			i--
		}
		if j >= 0 {
			v += b[j] - '0'
			j--
		}
		carry = '0' + (v-'0')/2
		v = '0' + (v-'0')%2
		bs = append([]byte{v}, bs...)
	}
	return string(bs)
}
