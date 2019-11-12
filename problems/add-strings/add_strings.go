package problem415

func addStrings(num1 string, num2 string) string {
	ans, l1, l2, carry := "", len(num1)-1, len(num2)-1, byte('0')
	for l1 >= 0 || l2 >= 0 || carry != '0' {
		v := carry
		if l1 >= 0 {
			v += num1[l1] - '0'
			l1--
		}
		if l2 >= 0 {
			v += num2[l2] - '0'
			l2--
		}
		carry = '0' + (v-'0')/10
		v = '0' + (v-'0')%10
		ans = string(v) + ans
	}
	return ans
}
