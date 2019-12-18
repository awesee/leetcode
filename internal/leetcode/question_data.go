package leetcode

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const testTpl = `package {{packageName}}

import "testing"

type testType struct {
	in   int
	want int
}

func Test{{funcName}}(t *testing.T) {
	tests := [...]testType{
		{
			in:   0,
			want: 0,
		},
	}
	for _, tt := range tests {
		got := 0
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
`

// QuestionDataType - leetcode.QuestionDataType
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
	QuestionID         string             `json:"questionId"`
	QuestionFrontendID string             `json:"questionFrontendId"`
	BoundTopicID       int                `json:"boundTopicId"`
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
	TopicTags          []TagType          `json:"topicTags"`
	CodeSnippets       []codeSnippetsType `json:"codeSnippets"`
	Hints              []string           `json:"hints"`
	MysqlSchemas       []string           `json:"mysqlSchemas"`
}

func (question *questionType) SaveContent() {
	if question.TitleSlug != "" {
		filePutContents(question.getFilePath("README.md"), question.getDescContent())
		question.saveMysqlSchemas()
	}
}

func (question *questionType) GetSimilarQuestion() (sq []similarQuestionType) {
	jsonDecode([]byte(question.SimilarQuestions), &sq)
	return
}

func (question *questionType) TitleSnake() string {
	return slugToSnake(question.TitleSlug)
}

func (question *questionType) LeetCodeURL() string {
	return "https://leetcode.com/problems/" + question.TitleSlug
}

func (question *questionType) PackageName() string {
	return fmt.Sprintf("problem%d", question.FrontendID())
}

func (question *questionType) SaveCodeSnippet() {
	if isLangMySQL(question.TitleSlug) {
		filePutContents(question.getFilePath(question.TitleSnake()+".sql"), []byte("# Write your MySQL query statement below\n"))
	}
	langSupport := [...]struct {
		slug   string
		handle func(*questionType, codeSnippetsType)
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
			return
		}
	}
}

func (question *questionType) getDescContent() []byte {
	var buf bytes.Buffer
	buf.WriteString(authInfo("description"))
	buf.WriteString(question.getNavigation())
	buf.WriteString(fmt.Sprintf("\n## [%d. %s%s](%s \"%s\")\n\n",
		question.FrontendID(),
		question.Title,
		question.Difficulty.Str(),
		question.LeetCodeURL(),
		question.translatedTitle(),
	))
	cts := filterContents(question.Content)
	// remove style
	reg := regexp.MustCompile(`<style[\S\s]+?</style>`)
	cts = reg.ReplaceAllString(cts, "")
	cts = strings.ReplaceAll(cts, "\n\n\t", "\n\t")
	cts = strings.TrimSpace(cts) + "\n"
	buf.WriteString(cts)
	buf.Write(question.getTopicTags())
	buf.Write(question.getSimilarQuestion())
	buf.Write(question.getHints())
	return buf.Bytes()
}

func (question *questionType) questionID() int {
	id, _ := strconv.Atoi(question.QuestionID)
	return id
}

func (question *questionType) FrontendID() int {
	id := titleSlugID[question.TitleSlug]
	if id != 0 {
		return id
	}
	id, _ = strconv.Atoi(question.QuestionFrontendID)
	return id
}

func (question *questionType) translatedTitle() string {
	if question.TranslatedTitle == "" {
		question.TranslatedTitle = translationSet[question.questionID()]
	}
	return question.TranslatedTitle
}

func (question *questionType) getNavigation() string {
	nav, pre, next := "\n%s\n%s\n%s\n", "< Previous", "Next >"
	problems := ProblemsAll().StatStatusPairs
	id := question.questionID()
	format := `[%s](../%s "%s")`
	for i, problem := range problems {
		if problem.Stat.QuestionID == id {
			if i < len(problems)-1 {
				pre = fmt.Sprintf(format, pre, problems[i+1].Stat.QuestionTitleSlug, problems[i+1].Stat.QuestionTitle)
			}
			if i > 0 {
				next = fmt.Sprintf(format, next, problems[i-1].Stat.QuestionTitleSlug, problems[i-1].Stat.QuestionTitle)
			}
			break
		}
	}
	return fmt.Sprintf(nav, pre, strings.Repeat("　", 16), next)
}

func (question *questionType) getTopicTags() []byte {
	tags := question.TopicTags
	var buf bytes.Buffer
	if len(tags) > 0 {
		buf.WriteString("\n### Related Topics\n")
	}
	format := "  [[%s](../../tag/%s/README.md)]\n"
	for _, tag := range tags {
		buf.WriteString(fmt.Sprintf(format, tag.Name, tag.Slug))
	}
	return buf.Bytes()
}

func (question *questionType) getHints() []byte {
	hints := question.Hints
	var buf bytes.Buffer
	if len(hints) > 0 {
		buf.WriteString("\n### Hints")
	}
	for i, hint := range hints {
		buf.WriteString(fmt.Sprintf("\n<details>\n<summary>Hint %d</summary>\n%s\n</details>\n", i+1, filterContents(hint)))
	}
	return buf.Bytes()
}

func (question *questionType) RenamePackageName() {
	packageName := fmt.Sprintf("package %s", question.PackageName())
	reg := regexp.MustCompile(`package \w+`)
	for _, ext := range [...]string{".go", "_test.go"} {
		cts := fileGetContents(question.getFilePath(question.TitleSnake() + ext))
		if len(cts) > 0 {
			content := reg.ReplaceAllString(string(cts), packageName)
			question.saveCodeContent(content, ext)
		}
	}
}

func (question *questionType) Restore() {
	filename := question.getFilePath("README.md")
	body := fileGetContents(filename)
	pattern := `## [^\n]+\n\n([\S\s]+)`
	if bytes.Contains(body, []byte("\n### ")) {
		pattern = `## [^\n]+\n\n([\S\s]+?)\n### `
	}
	reg := regexp.MustCompile(pattern)
	matches := reg.FindSubmatch(body)
	if len(matches) < 2 {
		return
	}
	question.Content = string(bytes.TrimSpace(matches[1]))
	name := fmt.Sprintf(questionDataFile, question.TitleSnake())
	filePutContents(getCachePath(name), jsonEncode(QuestionDataType{
		Data: dataType{Question: *question},
	}))
}

func (question *questionType) getSimilarQuestion() []byte {
	sq := question.GetSimilarQuestion()
	var buf bytes.Buffer
	if len(sq) > 0 {
		buf.WriteString("\n### Similar Questions\n")
	}
	format := "  1. [%s](../%s)%s\n"
	for _, q := range sq {
		buf.WriteString(fmt.Sprintf(format, q.Title, q.TitleSlug, q.Difficulty.Str()))
	}
	return buf.Bytes()
}

func (question *questionType) getFilePath(filename string) string {
	return filepath.Join("problems", question.TitleSlug, filename)
}

func (question *questionType) saveCodeContent(content, ext string, permX ...bool) {
	filePath := question.getFilePath(question.TitleSnake() + ext)
	filePutContents(filePath, []byte(content))
	if len(permX) > 0 && permX[0] {
		_ = os.Chmod(filePath, 0755)
	}
}

func (question *questionType) saveMysqlSchemas() {
	var buf bytes.Buffer
	for _, s := range question.MysqlSchemas {
		buf.WriteString(s + ";\n")
	}
	filePutContents(question.getFilePath("mysql_schemas.sql"), buf.Bytes())
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

// QuestionData - leetcode.QuestionData
func QuestionData(titleSlug string, isForce bool, graphQL ...string) (qd QuestionDataType) {
	jsonStr := `{
		"operationName": "questionData",
		"variables": {
			"titleSlug": "` + titleSlug + `"
		},
		"query": "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    __typename\n  }\n}\n"
	}`
	days := 3
	if isForce {
		days = 0
	}
	if len(graphQL) == 0 {
		graphQL = []string{graphQLCnURL}
	}
	name := fmt.Sprintf(questionDataFile, slugToSnake(titleSlug))
	filename := getCachePath(name)
	oldContent := getContent(filename)
	graphQLRequest(graphQL[0], jsonStr, name, days, &qd)
	if qd.Data.Question.Content == "" && oldContent != "" {
		qd.Data.Question.Content = oldContent
		filePutContents(filename, jsonEncode(qd))
	}
	if qd.Data.Question.TitleSlug == "" {
		if graphQL[0] == graphQLCnURL {
			return QuestionData(titleSlug, true, graphQLUrl)
		}
		for _, err := range qd.Errors {
			log.Println(titleSlug, err.Message)
		}
	}
	return
}

func handleCodeGolang(question *questionType, code codeSnippetsType) {
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

func handleCodeBash(question *questionType, code codeSnippetsType) {
	question.saveCodeContent("#!/usr/bin/env bash\n\n"+code.Code, ".bash", true)
}

func handleCodePython(question *questionType, code codeSnippetsType) {
	question.saveCodeContent("#!/usr/bin/env python\n\n"+code.Code, ".py", true)
}

func handleCodeSQL(question *questionType, code codeSnippetsType) {
	question.saveCodeContent(code.Code, ".sql")
	question.saveMysqlSchemas()
}

func filterContents(cts string) string {
	cts = strings.ReplaceAll(cts, "\r", "")
	cts = strings.ReplaceAll(cts, `src="/static`, `src="https://assets.leetcode.com/static_assets/public`)
	return cts
}

func getContent(filename string) string {
	var qd QuestionDataType
	cts := fileGetContents(filename)
	jsonDecode(cts, &qd)
	return qd.Data.Question.Content
}
