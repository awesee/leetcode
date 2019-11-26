// Package question implements the command question.
package question

import (
	"fmt"
	"strconv"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

// CmdQuestion - question.CmdQuestion
var CmdQuestion = &base.Command{
	Run:       runQuestion,
	UsageLine: "question <QuestionId>",
	Short:     "build problem solution file",
	Long:      "build problem's description,solution file.",
}

func runQuestion(cmd *base.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return
	}
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		if problem.Stat.FrontendQuestionID == id {
			fmt.Println(id, "\t"+problem.Stat.QuestionTitle)
			titleSlug := problem.Stat.QuestionTitleSlug
			question := leetcode.QuestionData(titleSlug, true).Data.Question
			if question.Content == "" && problem.PaidOnly == true && problem.Stat.QuestionArticleLive {
				question.Content = leetcode.GetDescription(problem.Stat.QuestionArticleSlug)
			}
			question.SaveContent()
			question.SaveCodeSnippet()
			return
		}
	}
}
