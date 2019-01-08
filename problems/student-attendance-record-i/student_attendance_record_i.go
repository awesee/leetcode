package student_attendance_record_i

import "strings"

func checkRecord(s string) bool {
	return !strings.Contains(s, "LLL") && strings.Count(s, "A") <= 1
}
