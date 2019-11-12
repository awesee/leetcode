package problem929

import "strings"

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	count := 0
	for _, email := range emails {
		plusIndex := strings.Index(email, "+")
		atIndex := strings.Index(email, "@")
		if plusIndex == -1 {
			plusIndex = atIndex
		}
		email = strings.ReplaceAll(email[0:plusIndex], ".", "") + email[atIndex:]
		if !m[email] {
			m[email] = true
			count++
		}
	}
	return count
}
