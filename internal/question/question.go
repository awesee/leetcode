package question

import (
	"strconv"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdQuestion = &base.Command{
	Run:       runQuestion,
	UsageLine: "question (QuestionFrontendId)",
	Short:     "build problem description file",
	Long:      "build problem's go,README.md file.",
}

func runQuestion(cmd *base.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
	}
	questionId, err := strconv.Atoi(args[0])
	base.CheckErr(err)
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		if problem.Stat.FrontendQuestionId == questionId {
			titleSlug := problem.Stat.QuestionTitleSlug
			question := leetcode.QuestionData(titleSlug).Data.Question
			question.SaveContent()
			question.SaveCodeSnippet()
		}
	}
}
