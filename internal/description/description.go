// Package description implements the command description.
package description

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

// CmdDescription - description.CmdDescription
var CmdDescription = &base.Command{
	Run:       runDescription,
	UsageLine: "description",
	Short:     "build all problems description file",
	Long:      "build all problems README.md file.",
	Hidden:    true,
}

func runDescription(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	var wg sync.WaitGroup
	jobs := make(chan leetcode.StatStatusPairsType)
	for i := 0; i < 1<<runtime.NumCPU(); i++ {
		go worker(jobs, &wg)
	}
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		fmt.Println(problem.Stat.FrontendQuestionID, "\t"+problem.Stat.QuestionTitle)
		wg.Add(1)
		jobs <- problem
	}
	wg.Wait()
}

func worker(jobs <-chan leetcode.StatStatusPairsType, wg *sync.WaitGroup) {
	for problem := range jobs {
		titleSlug := problem.Stat.QuestionTitleSlug
		question := leetcode.QuestionData(titleSlug, false).Data.Question
		if question.Content == "" && problem.PaidOnly == true && problem.Stat.QuestionArticleLive {
			question.Content = leetcode.GetDescription(problem.Stat.QuestionArticleSlug)
		}
		question.SaveContent()
		wg.Done()
	}
}
