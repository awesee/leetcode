package leetcode

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"
)

func QuestionData(titleSlug string) (qd questionDataType) {
	jsonStr := `{
		"operationName": "questionData",
		"variables": {
			"titleSlug": "` + titleSlug + `"
		},
		"query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"
	}`
	filename := "question_data_" + strings.Replace(titleSlug, "-", "_", -1) + ".json"
	graphQLRequest(filename, 30, jsonStr, &qd)
	return
}

type questionDataType struct {
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
	BoundTopicId       int                `json:"boundTopicId"`
	Title              string             `json:"title"`
	TitleSlug          string             `json:"titleSlug"`
	Content            string             `json:"content"`
	TranslatedTitle    string             `json:"translatedTitle"`
	TranslatedContent  string             `json:"translatedContent"`
	IsPaidOnly         paidType           `json:"isPaidOnly"`
	Difficulty         difficultyStrType  `json:"difficulty"`
	Likes              int                `json:"likes"`
	Dislikes           int                `json:"dislikes"`
	IsLiked            int                `json:"isLiked"`
	SimilarQuestions   string             `json:"similarQuestions"`
	TopicTags          []tagType          `json:"topicTags"`
	CodeSnippets       []codeSnippetsType `json:"codeSnippets"`
	Hints              []string           `json:"hints"`
	MysqlSchemas       []string           `json:"mysqlSchemas"`
}

type codeSnippetsType struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}

type similarQuestionType struct {
	Title           string            `json:"title"`
	TitleSlug       string            `json:"titleSlug"`
	Difficulty      difficultyStrType `json:"difficulty"`
	TranslatedTitle string            `json:"translatedTitle"`
}

type difficultyStrType string

func (d difficultyStrType) Str() (s string) {
	if d != "" {
		s = fmt.Sprintf(" (%s)", d)
	}
	return
}

func (question questionType) SaveContent() {
	fmt.Println(question.QuestionFrontendId, "\t"+question.Title, "saving...")
	filePutContents(question.getFilePath("README.md"), question.getDescContent())
}

func (question questionType) getDescContent() []byte {
	var buf bytes.Buffer
	buf.WriteString(authInfo("description"))
	buf.WriteString(fmt.Sprintf("\n## %s. %s%s\n\n", question.QuestionFrontendId, question.Title, question.Difficulty.Str()))
	content := strings.Replace(question.Content, "\r", "", -1)
	// remove style
	reg := regexp.MustCompile(`<style[\S\s]+?</style>`)
	content = reg.ReplaceAllString(content, "")
	content = strings.Replace(content, "\n\n\t", "\n\t", -1)
	content = strings.TrimSpace(content)
	buf.WriteString(content + "\n")
	buf.Write(question.getTopicTags())
	buf.Write(question.getSimilarQuestion())
	buf.Write(question.getHints())
	return buf.Bytes()
}

func (question questionType) getTopicTags() []byte {
	tags := question.TopicTags
	var buf bytes.Buffer
	if len(tags) > 0 {
		buf.WriteString("\n### Related Topics\n")
	}
	format := "  [[%s](https://github.com/openset/leetcode/tree/master/tag/%s/README.md)]\n"
	for _, tag := range tags {
		buf.WriteString(fmt.Sprintf(format, tag.Name, tag.Slug))
	}
	return buf.Bytes()
}

func (question questionType) getSimilarQuestion() []byte {
	var sq []similarQuestionType
	jsonDecode([]byte(question.SimilarQuestions), &sq)
	var buf bytes.Buffer
	if len(sq) > 0 {
		buf.WriteString("\n### Similar Questions\n")
	}
	format := "  1. [%s](https://github.com/openset/leetcode/tree/master/problems/%s)%s\n"
	for _, q := range sq {
		buf.WriteString(fmt.Sprintf(format, q.Title, q.TitleSlug, q.Difficulty.Str()))
	}
	return buf.Bytes()
}

func (question questionType) getHints() []byte {
	hints := question.Hints
	var buf bytes.Buffer
	if len(hints) > 0 {
		buf.WriteString("\n### Hints\n")
	}
	for i, hint := range hints {
		hint = strings.Replace(hint, "\r", "", -1)
		buf.WriteString(fmt.Sprintf("<details>\n<summary>Hint %d</summary>\n%s\n</details>\n", i+1, hint))
	}
	return buf.Bytes()
}

func (question questionType) getFilePath(filename string) string {
	return path.Join("problems", question.TitleSlug, filename)
}

func (question questionType) TitleSnake() string {
	return strings.Replace(question.TitleSlug, "-", "_", -1)
}

func (question questionType) PackageName() string {
	snake := question.TitleSnake()
	if snake != "" && unicode.IsNumber(rune(snake[0])) {
		snake = "p_" + snake
	}
	return snake
}

func (question questionType) SaveCodeSnippet() {
	langSupport := [...]struct {
		slug   string
		handle func(questionType, codeSnippetsType)
	}{
		{"golang", handleCodeGolang},
		{"python3", handleCodePython},
		{"python", handleCodePython},
		{"bash", handleCodeBash},
		{"mysql", handleCodeSQL},
		{"mssql", handleCodeSQL},
		{"oraclesql", handleCodeSQL},
	}
	codeSet := make(map[string]codeSnippetsType)
	for _, code := range question.CodeSnippets {
		codeSet[code.LangSlug] = code
	}
	for _, lang := range langSupport {
		if code, ok := codeSet[lang.slug]; ok {
			lang.handle(question, code)
			break
		}
	}
}

func (question questionType) saveCodeContent(content, ext string, permX ...bool) {
	filePath := question.getFilePath(question.TitleSnake() + ext)
	filePutContents(filePath, []byte(content))
	if len(permX) > 0 && permX[0] == true {
		os.Chmod(filePath, 0755)
	}
}

func handleCodeGolang(question questionType, code codeSnippetsType) {
	content := fmt.Sprintf("package %s\n\n", question.PackageName())
	content += code.Code + "\n"
	question.saveCodeContent(content, ".go")
	testExt := "_test.go"
	contents := fileGetContents(question.getFilePath(question.TitleSnake() + testExt))
	if bytes.Count(contents, []byte("\n")) <= 2 {
		// match function name
		reg := regexp.MustCompile(`func (\w+?)\(`)
		matches := reg.FindStringSubmatch(code.Code)
		funcName := "Func"
		if len(matches) >= 2 {
			funcName = matches[1]
		}
		content = strings.NewReplacer(
			"{{packageName}}", question.PackageName(),
			"{{funcName}}", strings.Title(funcName),
		).Replace(testTpl)
		question.saveCodeContent(content, testExt)
	}
}

func handleCodeBash(question questionType, code codeSnippetsType) {
	question.saveCodeContent("#!/usr/bin/env bash\n\n"+code.Code, ".bash", true)
}

func handleCodePython(question questionType, code codeSnippetsType) {
	question.saveCodeContent("#!/usr/bin/env python\n\n"+code.Code, ".py", true)
}

func handleCodeSQL(question questionType, code codeSnippetsType) {
	question.saveCodeContent(code.Code, ".sql")
	var buf bytes.Buffer
	for _, s := range question.MysqlSchemas {
		buf.WriteString(s + ";\n")
	}
	filePutContents(question.getFilePath("mysql_schemas.sql"), buf.Bytes())
}

const testTpl = `package {{packageName}}

import "testing"

type caseType struct {
	input    int
	expected int
}

func Test{{funcName}}(t *testing.T) {
	tests := [...]caseType{
		{
			input:    0,
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := 0
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
`
