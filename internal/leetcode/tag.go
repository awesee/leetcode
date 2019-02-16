package leetcode

import (
	"bytes"
	"fmt"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/openset/leetcode/internal/client"
)

var initTags []tagType

func init() {
	html := remember(problemsetAllFile, 7, func() []byte {
		return client.Get(problemsetAllUrl)
	})
	reg := regexp.MustCompile(`href="/tag/(\S+?)/"`)
	for _, matches := range reg.FindAllStringSubmatch(string(html), -1) {
		if len(matches) >= 2 {
			initTags = append(initTags, tagType{Slug: matches[1]})
		}
	}
}

func GetTags() (tags []tagType) {
	cts := fileGetContents(tagsFile)
	jsonDecode(cts, &tags)
	tags = tagsUnique(tags)
	return
}

func saveTags(tags []tagType) {
	tags = append(GetTags(), tags...)
	filePutContents(tagsFile, jsonEncode(tagsUnique(tags)))
}

func tagsUnique(tags []tagType) []tagType {
	rs, top := make([]tagType, 0, len(tags)), 1
	tags = append(initTags, tags...)
	var flag = make(map[string]int)
	for _, tag := range tags {
		i := flag[tag.Slug]
		if i == 0 {
			rs = append(rs, tag)
			flag[tag.Slug] = top
			top++
		} else {
			if tag.Name != "" {
				rs[i-1].Name = tag.Name
			}
			if tag.TranslatedName != "" {
				rs[i-1].TranslatedName = tag.TranslatedName
			}
		}
	}
	return rs
}

func GetTopicTag(slug string) (tt topicTagType) {
	jsonStr := `{
		"operationName": "getTopicTag",
		"variables": {
			"slug": "` + slug + `"
		},
		"query": "query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    translatedName\n    questions {\n      status\n      questionId\n      questionFrontendId\n      title\n      titleSlug\n      translatedTitle\n      stats\n      difficulty\n      isPaidOnly\n      topicTags {\n        name\n        translatedName\n        slug\n        __typename\n      }\n      companyTags {\n        name\n        translatedName\n        slug\n        __typename\n      }\n      __typename\n    }\n    frequencies\n    __typename\n  }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n      __typename\n    }\n    privateFavorites {\n      ...favoriteFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n    __typename\n  }\n  __typename\n}\n"
	}`
	filename := "topic_tag_" + strings.Replace(slug, "-", "_", -1) + ".json"
	graphQLRequest(filename, 3, jsonStr, &tt)
	return
}

type tagType struct {
	Name           string
	Slug           string
	TranslatedName string
}

type topicTagType struct {
	Errors []errorType `json:"errors"`
	Data   ttDataType  `json:"data"`
}

type ttDataType struct {
	TopicTag ttType `json:"topicTag"`
}

type ttType struct {
	Name           string           `json:"name"`
	TranslatedName string           `json:"translatedName"`
	Questions      []ttQuestionType `json:"questions"`
}

type ttQuestionType struct {
	QuestionId         string    `json:"questionId"`
	QuestionFrontendId string    `json:"questionFrontendId"`
	Title              string    `json:"title"`
	TitleSlug          string    `json:"titleSlug"`
	TranslatedTitle    string    `json:"translatedTitle"`
	TranslatedContent  string    `json:"translatedContent"`
	IsPaidOnly         paidType  `json:"isPaidOnly"`
	Difficulty         string    `json:"difficulty"`
	TopicTags          []tagType `json:"topicTags"`
}

func (question ttQuestionType) TagsStr() string {
	var buf bytes.Buffer
	format := "[[%s](https://github.com/openset/leetcode/tree/master/tag/%s/README.md)] "
	for _, tag := range question.TopicTags {
		buf.WriteString(fmt.Sprintf(format, tag.ShowName(), tag.Slug))
	}
	saveTags(question.TopicTags)
	return string(buf.Bytes())
}

func (tag tagType) SaveContents() {
	questions := GetTopicTag(tag.Slug).Data.TopicTag.Questions
	sort.Slice(questions, func(i, j int) bool {
		m, _ := strconv.Atoi(questions[i].QuestionFrontendId)
		n, _ := strconv.Atoi(questions[j].QuestionFrontendId)
		return m > n
	})
	var buf bytes.Buffer
	buf.WriteString(authInfo("tag"))
	buf.WriteString(fmt.Sprintf("\n## %s\n\n", tag.ShowName()))
	buf.WriteString("| # | 题名 | 标签 | 难度 |\n")
	buf.WriteString("| :-: | - | - | :-: |\n")
	format := "| %s | [%s](https://github.com/openset/leetcode/tree/master/problems/%s)%s | %s | %s |\n"
	for _, question := range questions {
		if question.TranslatedTitle == "" {
			question.TranslatedTitle = question.Title
		}
		buf.WriteString(fmt.Sprintf(format, question.QuestionFrontendId, question.TranslatedTitle, question.TitleSlug, question.IsPaidOnly.Str(), question.TagsStr(), question.Difficulty))
	}
	filename := path.Join("tag", tag.Slug, "README.md")
	filePutContents(filename, buf.Bytes())
}

func (tag tagType) ShowName() string {
	if tag.TranslatedName != "" {
		return tag.TranslatedName
	}
	return tag.Name
}
