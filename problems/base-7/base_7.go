package problem504

func convertToBase7(num int) string {
	ans := ""
	if num < 0 {
		ans, num = "-", -num
	}
	if num >= 7 {
		ans += convertToBase7(num / 7)
	}
	return ans + string('0'+num%7)
}
