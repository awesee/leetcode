package leetcode

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/client"
)

var (
	initTags []TagType
	tagsFile = filepath.Join("tag", "tags.json")
)

// TagType - leetcode.TagType
type TagType struct {
	Name           string
	Slug           string
	TranslatedName string
}

// SaveContents - leetcode.SaveContents
func (tag *TagType) SaveContents() {
	questions := GetTopicTag(tag.Slug).Data.TopicTag.Questions
	sort.Slice(questions, func(i, j int) bool {
		return questions[i].frontendID() > questions[j].frontendID()
	})
	var buf bytes.Buffer
	buf.WriteString(authInfo("tag"))
	buf.WriteString(fmt.Sprintf("\n## [话题分类](../README.md) > %s\n\n", tag.name()))
	buf.WriteString("| # | 题目 | 标签 | 难度 |\n")
	buf.WriteString("| :-: | - | - | :-: |\n")
	format := "| %d | [%s](../../problems/%s)%s | %s | %s |\n"
	for _, question := range questions {
		if question.TranslatedTitle == "" {
			question.TranslatedTitle = question.Title
		}
		if question.frontendID() > 0 {
			buf.WriteString(fmt.Sprintf(format, question.frontendID(), question.TranslatedTitle, question.TitleSlug, question.IsPaidOnly.Str(), question.TagsStr(), question.Difficulty))
		}
	}
	filename := filepath.Join("tag", tag.Slug, "README.md")
	filePutContents(filename, buf.Bytes())
}

func (tag *TagType) name() string {
	if tag.TranslatedName != "" {
		return tag.TranslatedName
	}
	return tag.Name
}

// TopicTagType - leetcode.TopicTagType
type TopicTagType struct {
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
	QuestionID         string    `json:"questionId"`
	QuestionFrontendID string    `json:"questionFrontendId"`
	Title              string    `json:"title"`
	TitleSlug          string    `json:"titleSlug"`
	TranslatedTitle    string    `json:"translatedTitle"`
	TranslatedContent  string    `json:"translatedContent"`
	IsPaidOnly         paidType  `json:"isPaidOnly"`
	Difficulty         string    `json:"difficulty"`
	TopicTags          []TagType `json:"topicTags"`
}

func (question *ttQuestionType) TagsStr() string {
	var buf bytes.Buffer
	format := "[[%s](../%s/README.md)] "
	for _, tag := range question.TopicTags {
		buf.WriteString(fmt.Sprintf(format, tag.name(), tag.Slug))
	}
	saveTags(question.TopicTags)
	return buf.String()
}

func (question *ttQuestionType) frontendID() int {
	id := titleSlugID[question.TitleSlug]
	if id != 0 {
		return id
	}
	id, _ = strconv.Atoi(question.QuestionFrontendID)
	return id
}

// GetTags - leetcode.GetTags
func GetTags() (tags []TagType) {
	cts := fileGetContents(tagsFile)
	jsonDecode(cts, &tags)
	tags = tagsUnique(tags)
	return
}

// GetTopicTag - leetcode.GetTopicTag
func GetTopicTag(slug string) (tt TopicTagType) {
	jsonStr := `{
		"operationName": "getTopicTag",
		"variables": {
			"slug": "` + slug + `"
		},
		"query": "query getTopicTag($slug: String!) {\n  topicTag(slug: $slug) {\n    name\n    translatedName\n    questions {\n      status\n      questionId\n      questionFrontendId\n      title\n      titleSlug\n      translatedTitle\n      stats\n      difficulty\n      isPaidOnly\n      topicTags {\n        name\n        translatedName\n        slug\n        __typename\n      }\n      __typename\n    }\n    frequencies\n    __typename\n  }\n  favoritesLists {\n    publicFavorites {\n      ...favoriteFields\n      __typename\n    }\n    privateFavorites {\n      ...favoriteFields\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment favoriteFields on FavoriteNode {\n  idHash\n  id\n  name\n  isPublicFavorite\n  viewCount\n  creator\n  isWatched\n  questions {\n    questionId\n    title\n    titleSlug\n    __typename\n  }\n  __typename\n}\n"
	}`
	filename := fmt.Sprintf(topicTagFile, slugToSnake(slug))
	graphQLRequest(graphQLCnURL, jsonStr, filename, 2, &tt)
	if tt.Data.TopicTag.Name == "" {
		_ = os.Remove(getCachePath(filename))
		for _, err := range tt.Errors {
			log.Println(slug, err.Message)
		}
	}
	return
}

func saveTags(tags []TagType) {
	base.Mutex.Lock()
	tags = append(GetTags(), tags...)
	filePutContents(tagsFile, jsonEncode(tagsUnique(tags)))
	base.Mutex.Unlock()
}

func tagsUnique(tags []TagType) []TagType {
	rs, top := make([]TagType, 0, len(tags)), 1
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
				rs[i-1].TranslatedName = tag.Name
			}
			if tag.TranslatedName != "" {
				rs[i-1].TranslatedName = tag.TranslatedName
			}
		}
	}
	return rs
}

func init() {
	html := remember(problemsetAllFile, 3, func() []byte {
		return client.Get(problemsetAllURL)
	})
	reg := regexp.MustCompile(`href="/tag/(\S+?)/"`)
	for _, matches := range reg.FindAllStringSubmatch(string(html), -1) {
		if len(matches) >= 2 {
			initTags = append(initTags, TagType{Slug: matches[1]})
		}
	}
}
