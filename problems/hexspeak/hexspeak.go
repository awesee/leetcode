package problem1271

import (
	"fmt"
	"strconv"
	"strings"
)

func toHexspeak(num string) string {
	n, _ := strconv.Atoi(num)
	hex := fmt.Sprintf("%X", n)
	hex = strings.NewReplacer("0", "O", "1", "I").Replace(hex)
	for _, v := range hex {
		if v < 'A' {
			return "ERROR"
		}
	}
	return hex
}
