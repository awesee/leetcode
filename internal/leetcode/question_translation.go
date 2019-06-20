package leetcode

import (
	"log"
	"os"
	"strconv"
)

func GetQuestionTranslation() (qt questionTranslationType) {
	jsonStr := `{
		"operationName": "getQuestionTranslation",
		"variables": {},
		"query": "query getQuestionTranslation($lang: String) {\n  translations: allAppliedQuestionTranslations(lang: $lang) {\n    title\n    question {\n      questionId\n      __typename\n    }\n    __typename\n  }\n}\n"
	}`
	graphQLRequest(graphQLCnUrl, jsonStr, questionTranslationFile, 2, &qt)
	if qt.Data.Translations == nil {
		_ = os.Remove(getCachePath(questionTranslationFile))
		for _, err := range qt.Errors {
			log.Println(err.Message)
		}
	}
	return
}

type questionTranslationType struct {
	Errors []errorType `json:"errors"`
	Data   qtDataType  `json:"data"`
}

type qtDataType struct {
	Translations []translationsType `json:"translations"`
}

type translationsType struct {
	Title    string         `json:"title"`
	Question questionIdType `json:"question"`
	TypeName string         `json:"__typename"`
}

type questionIdType struct {
	QuestionId string `json:"questionId"`
	TypeName   string `json:"__typename"`
}

func init() {
	translation := GetQuestionTranslation()
	for _, item := range translation.Data.Translations {
		questionId := item.Question.QuestionId
		if id, err := strconv.Atoi(questionId); err == nil {
			translationSet[id] = item.Title
		}
	}
}
