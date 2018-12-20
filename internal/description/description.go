package description

import (
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
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		titleSlug := problem.Stat.QuestionTitleSlug
		leetcode.QuestionData(titleSlug).Data.Question.SaveContent()
	}
}
