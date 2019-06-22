package description

import (
	"fmt"
	"sync"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdDescription = &base.Command{
	Run:       runDescription,
	UsageLine: "description",
	Short:     "build all problems description file",
	Long:      "build all problems README.md file.",
}

func runDescription(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	var wg sync.WaitGroup
	tokens := make(chan bool, 1<<7)
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		wg.Add(1)
		tokens <- true
		fmt.Println(problem.Stat.FrontendQuestionId, "\t"+problem.Stat.QuestionTitle)
		go func(problem leetcode.StatStatusPairsType) {
			titleSlug := problem.Stat.QuestionTitleSlug
			question := leetcode.QuestionData(titleSlug, false).Data.Question
			if question.Content == "" && problem.PaidOnly == true && problem.Stat.QuestionArticleLive {
				question.Content = leetcode.GetDescription(problem.Stat.QuestionArticleSlug)
			}
			question.SaveContent()
			<-tokens
			wg.Done()
		}(problem)
	}
	wg.Wait()
}
