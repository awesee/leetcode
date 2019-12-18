// Package open implements the command open.
package open

import (
	"fmt"
	"strconv"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/browser"
	"github.com/openset/leetcode/internal/leetcode"
)

// CmdOpen - open.CmdOpen
var CmdOpen = &base.Command{
	Run:       runOpen,
	UsageLine: "open <QuestionId>",
	Short:     "open a problem solution in browser",
	Long:      "open a problem solution in browser.",
}

func runOpen(cmd *base.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		cmd.Usage()
		return
	}
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		if problem.Stat.FrontendQuestionID == id {
			titleSlug := problem.Stat.QuestionTitleSlug
			browser.Open(fmt.Sprintf("%s/problems/%s", base.URL, titleSlug))
			break
		}
	}
}
