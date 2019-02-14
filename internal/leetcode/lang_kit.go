package leetcode

var langSet = make(map[string]string)

func saveAsMySQL(titleSlug string) {
	langSet[titleSlug] = "MySQL"
}

func isLangMySQL(titleSlug string) bool {
	return langSet[titleSlug] == "MySQL"
}

func saveAsBash(titleSlug string) {
	langSet[titleSlug] = "Bash"
}

func getLang(titleSlug string) string {
	if lang, ok := langSet[titleSlug]; ok {
		return lang
	}
	return "Go"
}
