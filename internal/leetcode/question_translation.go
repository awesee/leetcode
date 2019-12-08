package leetcode

import (
	"log"
	"os"
	"strconv"
)

// QuestionTranslationType - leetcode.QuestionTranslationType
type QuestionTranslationType struct {
	Errors []errorType `json:"errors"`
	Data   qtDataType  `json:"data"`
}

type qtDataType struct {
	Translations []translationsType `json:"translations"`
}

type translationsType struct {
	Title      string `json:"title"`
	QuestionID string `json:"questionId"`
	TypeName   string `json:"__typename"`
}

// GetQuestionTranslation - leetcode.GetQuestionTranslation
func GetQuestionTranslation() (qt QuestionTranslationType) {
	jsonStr := `{
		"operationName": "getQuestionTranslation",
		"variables": {},
		"query": "query getQuestionTranslation($lang: String) {\n  translations: allAppliedQuestionTranslations(lang: $lang) {\n    title\n    questionId\n    __typename\n  }\n}\n"
	}`
	graphQLRequest(graphQLCnURL, jsonStr, questionTranslationFile, 2, &qt)
	if qt.Data.Translations == nil {
		_ = os.Remove(getCachePath(questionTranslationFile))
		for _, err := range qt.Errors {
			log.Println(err.Message)
		}
	}
	return
}

func init() {
	translation := GetQuestionTranslation()
	for _, item := range translation.Data.Translations {
		id := item.QuestionID
		if id, err := strconv.Atoi(id); err == nil {
			translationSet[id] = item.Title
		}
	}
}
