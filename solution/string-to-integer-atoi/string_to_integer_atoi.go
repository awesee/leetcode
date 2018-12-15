package string_to_integer_atoi

import (
	"strings"
	"unicode"
	"strconv"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	prefix := ""
	if strings.HasPrefix(str, "-") || strings.HasPrefix(str, "+") {
		prefix = str[0:1]
		str = str[1:]
	}
	i := strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
	if i > -1 {
		str = str[0:i]
	}
	r, _ := strconv.Atoi(prefix + str)
	if r < -1<<31 {
		r = -1 << 31
	} else if r > 1<<31-1 {
		r = 1<<31 - 1
	}
	return r
}
