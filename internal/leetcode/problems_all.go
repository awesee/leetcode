package leetcode

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/openset/leetcode/internal/client"
)

// ProblemsType - leetcode.ProblemsType
type ProblemsType struct {
	UserName        string                `json:"user_name"`
	NumSolved       int                   `json:"num_solved"`
	NumTotal        int                   `json:"num_total"`
	AcEasy          int                   `json:"ac_easy"`
	AcMedium        int                   `json:"ac_medium"`
	AcHard          int                   `json:"ac_hard"`
	StatStatusPairs []StatStatusPairsType `json:"stat_status_pairs"`
	FrequencyHigh   int                   `json:"frequency_high"`
	FrequencyMid    int                   `json:"frequency_mid"`
	CategorySlug    string                `json:"category_slug"`
}

// StatStatusPairsType - leetcode.StatStatusPairsType
type StatStatusPairsType struct {
	Stat       statType       `json:"stat"`
	Status     string         `json:"status"`
	Difficulty difficultyType `json:"difficulty"`
	PaidOnly   paidType       `json:"paid_only"`
	IsFavor    bool           `json:"is_favor"`
	Frequency  int            `json:"frequency"`
	Progress   int            `json:"progress"`
}

// WriteRow - leetcode.WriteRow
func (problem *StatStatusPairsType) WriteRow(buf *bytes.Buffer, path string) {
	format := "| <span id=\"%d\">%d</span> | [%s](https://leetcode.com/problems/%s%s)%s | [%s](%s/%s) | %s |\n"
	id := problem.Stat.FrontendQuestionID
	stat := problem.Stat
	title := strings.TrimSpace(problem.Stat.QuestionTitle)
	titleSlug := stat.QuestionTitleSlug
	levelName := problem.Difficulty.LevelName()
	buf.WriteString(fmt.Sprintf(format, id, id,
		title, titleSlug, stat.TranslationTitle(), problem.PaidOnly.Str(),
		stat.Lang(), path, titleSlug,
		levelName,
	))
}

type statType struct {
	QuestionID          int    `json:"question_id"`
	QuestionArticleLive bool   `json:"question__article__live"`
	QuestionArticleSlug string `json:"question__article__slug"`
	QuestionTitle       string `json:"question__title"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionHide        bool   `json:"question__hide"`
	TotalAcs            int    `json:"total_acs"`
	TotalSubmitted      int    `json:"total_submitted"`
	FrontendQuestionID  int    `json:"frontend_question_id"`
	IsNewQuestion       bool   `json:"is_new_question"`
}

func (s *statType) Lang() string {
	return getLang(s.QuestionTitleSlug)
}

func (s *statType) QuestionTitleSnake() string {
	return slugToSnake(s.QuestionTitleSlug)
}

func (s *statType) TranslationTitle() string {
	title := translationSet[s.QuestionID]
	if title != "" {
		title = fmt.Sprintf(` "%s"`, title)
	}
	return title
}

type difficultyType struct {
	Level int `json:"level"`
}

func (d *difficultyType) LevelName() string {
	m := map[int]string{
		1: "Easy",
		2: "Medium",
		3: "Hard",
	}
	return m[d.Level]
}

type paidType bool

func (p paidType) Str() string {
	if p {
		return LockStr
	}
	return ""
}

// ProblemsAll - leetcode.ProblemsAll
func ProblemsAll() (ps ProblemsType) {
	data := remember(problemsAllFile, 2, func() []byte {
		return client.Get(apiProblemsAllURL)
	})
	jsonDecode(data, &ps)
	sort.Slice(ps.StatStatusPairs, func(i, j int) bool {
		return ps.StatStatusPairs[i].Stat.FrontendQuestionID > ps.StatStatusPairs[j].Stat.FrontendQuestionID
	})
	return
}

func init() {
	questions := ProblemsAll().StatStatusPairs
	for _, item := range questions {
		titleSlugID[item.Stat.QuestionTitleSlug] = item.Stat.FrontendQuestionID
	}
}
