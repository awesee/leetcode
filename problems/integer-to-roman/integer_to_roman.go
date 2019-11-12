package problem12

func intToRoman(num int) string {
	base := [...]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	bs := [...]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ans := ""
	for i, b := range base {
		for num/b > 0 {
			ans += bs[i]
			num -= b
		}
	}
	return ans
}
