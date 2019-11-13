// Package test implements the command test.
package test

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

// CmdTest - test.CmdTest
var CmdTest = &base.Command{
	Run:       runTest,
	UsageLine: "test [QuestionId]",
	Short:     "run go test",
	Long:      "automates testing the packages.",
}

func runTest(cmd *base.Command, args []string) {
	if len(args) > 1 {
		cmd.Usage()
		return
	}
	if _, err := os.Stat("problems"); err == nil {
		target := "./..."
		if len(args) == 1 {
			if id, err := strconv.Atoi(args[0]); err == nil {
				problems := leetcode.ProblemsAll()
				for _, problem := range problems.StatStatusPairs {
					if problem.Stat.FrontendQuestionID == id {
						target = "./problems/" + problem.Stat.QuestionTitleSlug
						break
					}
				}
			}
		}
		c := exec.Command("go", "test", target)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		err := c.Run()
		base.CheckErr(err)
	}
}
