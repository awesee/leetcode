package problem65

import (
	"strconv"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	_, err := strconv.ParseFloat(s, 64)
	return err == nil || err.(*strconv.NumError).Err != strconv.ErrSyntax
}
