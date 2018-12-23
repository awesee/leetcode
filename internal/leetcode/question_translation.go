package leetcode

func GetQuestionTranslation() (qt questionTranslationType) {
	jsonStr := `{
	    "operationName": "getQuestionTranslation",
	    "variables": {},
	    "query": "query getQuestionTranslation($lang: String) {\n  translations: allAppliedQuestionTranslations(lang: $lang) {\n    title\n    question {\n      questionId\n      __typename\n    }\n    __typename\n  }\n}\n"
	}`
	graphQLRequest(questionTranslationFile, jsonStr, &qt)
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
