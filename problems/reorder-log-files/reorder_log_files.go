package reorder_log_files

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	var logSort logType = logs
	sort.Stable(logSort)
	return logSort
}

type logType []string

func (l logType) Len() int {
	return len(l)
}

func (l logType) Less(i, j int) bool {
	s1 := strings.SplitN(l[i], " ", 2)
	s2 := strings.SplitN(l[j], " ", 2)
	f1, f2 := "0"+s1[1], "0"+s2[1]
	if unicode.IsNumber(rune(f1[1])) {
		f1 = "1"
	}
	if unicode.IsNumber(rune(f2[1])) {
		f2 = "1"
	}
	return f1 < f2
}

func (l logType) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
