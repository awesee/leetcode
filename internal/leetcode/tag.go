package leetcode

func GetTags() (tags []tagsType) {
	data := fileGetContents("tag/tags.json")
	jsonDecode(data, &tags)
	return
}

type tagsType struct {
	Title           string
	TitleSlug       string
	TranslatedTitle string
}
