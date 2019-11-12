package problem412

import "strconv"

func fizzBuzz(n int) []string {
	s := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		str := ""
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str = strconv.Itoa(i)
		}
		s = append(s, str)
	}
	return s
}
