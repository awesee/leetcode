package leetcode

import (
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
)

func QuestionData(titleSlug string) (qd QuestionDataType) {
	jsonStr := `{
  		"operationName": "questionData",
  		"variables": {
    		"titleSlug": "` + titleSlug + `"
  		},
  		"query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"
	}`
	resp, err := http.Post(GraphqlUrl, "application/json", strings.NewReader(jsonStr))
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	err = json.Unmarshal(body, &qd)
	checkErr(err)
	return
}

type QuestionDataType struct {
	Errors []errorType `json:"errors"`
	Data   dataType    `json:"data"`
}

type errorType struct {
	Message string `json:"message"`
}

type dataType struct {
	Question questionType `json:"question"`
}

type questionType struct {
	QuestionId         string             `json:"questionId"`
	QuestionFrontendId string             `json:"questionFrontendId"`
	BoundTopicId       string             `json:"boundTopicId"`
	Title              string             `json:"title"`
	TitleSlug          string             `json:"titleSlug"`
	Content            string             `json:"content"`
	TranslatedTitle    string             `json:"translatedTitle"`
	TranslatedContent  string             `json:"translatedContent"`
	IsPaidOnly         bool               `json:"isPaidOnly"`
	Difficulty         string             `json:"difficulty"`
	Likes              int                `json:"likes"`
	Dislikes           int                `json:"dislikes"`
	IsLiked            int                `json:"isLiked"`
	SimilarQuestions   string             `json:"similarQuestions"`
	CodeSnippets       []codeSnippetsType `json:"codeSnippets"`
}

type codeSnippetsType struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

type SimilarQuestionType struct {
	Title           string `json:"title"`
	TitleSlug       string `json:"titleSlug"`
	Difficulty      string `json:"difficulty"`
	TranslatedTitle string `json:"translatedTitle"`
}
