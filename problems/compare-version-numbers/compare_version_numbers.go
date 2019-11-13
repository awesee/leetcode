package problem165

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")
	i, l1 := 0, len(v1s)
	j, l2 := 0, len(v2s)
	for i < l1 || j < l2 {
		v1, v2 := "0", "0"
		if i < l1 {
			v1 = v1s[i]
			i++
		}
		if j < l2 {
			v2 = v2s[j]
			j++
		}
		n1, _ := strconv.Atoi(v1)
		n2, _ := strconv.Atoi(v2)
		if n1 != n2 {
			if n1 < n2 {
				return -1
			}
			return 1
		}
	}
	return 0
}
