package problem8

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	prefix := ""
	if strings.HasPrefix(str, "-") || strings.HasPrefix(str, "+") {
		prefix = str[0:1]
		str = str[1:]
	}
	i := strings.IndexFunc(str, func(r rune) bool {
		return !unicode.IsDigit(r)
	})
	if i > -1 {
		str = str[0:i]
	}
	r, _ := strconv.Atoi(prefix + str)
	if r < math.MinInt32 {
		r = math.MinInt32
	} else if r > math.MaxInt32 {
		r = math.MaxInt32
	}
	return r
}
