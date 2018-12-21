package leetcode

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type ProblemsAllType struct {
	UserName        string                `json:"user_name"`
	NumSolved       int                   `json:"num_solved"`
	NumTotal        int                   `json:"num_total"`
	AcEasy          int                   `json:"ac_easy"`
	AcMedium        int                   `json:"ac_medium"`
	AcHard          int                   `json:"ac_hard"`
	StatStatusPairs []statStatusPairsType `json:"stat_status_pairs"`
	FrequencyHigh   int                   `json:"frequency_high"`
	FrequencyMid    int                   `json:"frequency_mid"`
	CategorySlug    string                `json:"category_slug"`
}

type statStatusPairsType struct {
	Stat       statType       `json:"stat"`
	Status     string         `json:"status"`
	Difficulty difficultyType `json:"difficulty"`
	PaidOnly   bool           `json:"paid_only"`
	IsFavor    bool           `json:"is_favor"`
	Frequency  int            `json:"frequency"`
	Progress   int            `json:"progress"`
}

type statType struct {
	QuestionId          int    `json:"question_id"`
	QuestionArticleLive bool   `json:"question__article__live"`
	QuestionArticleSlug string `json:"question__article__slug"`
	QuestionTitle       string `json:"question__title"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionHide        bool   `json:"question__hide"`
	TotalAcs            int    `json:"total_acs"`
	TotalSubmitted      int    `json:"total_submitted"`
	FrontendQuestionId  int    `json:"frontend_question_id"`
	IsNewQuestion       bool   `json:"is_new_question"`
}

type difficultyType struct {
	Level int `json:"level"`
}

func (s statType) QuestionTitleSnake() string {
	return strings.Replace(s.QuestionTitleSlug, "-", "_", -1)
}

func (d difficultyType) LevelName() string {
	m := map[int]string{
		1: "Easy",
		2: "Medium",
		3: "Hard",
	}
	return m[d.Level]
}

func ProblemsAll() (pa ProblemsAllType) {
	resp, err := http.Get(ApiProblemsAllUrl)
	checkErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	var buf bytes.Buffer
	err = json.Indent(&buf, body, "", "\t")
	checkErr(err)
	filePutContents(getCachePath(problemsAllFile), buf.Bytes())
	err = json.Unmarshal(body, &pa)
	checkErr(err)
	return
}
