package problem67

func addBinary(a string, b string) string {
	ans, l1, l2, carry := "", len(a)-1, len(b)-1, byte('0')
	for l1 >= 0 || l2 >= 0 || carry != '0' {
		v := carry
		if l1 >= 0 {
			v += a[l1] - '0'
			l1--
		}
		if l2 >= 0 {
			v += b[l2] - '0'
			l2--
		}
		carry = '0' + (v-'0')/2
		v = '0' + (v-'0')%2
		ans = string(v) + ans
	}
	return ans
}
