package page

import (
	"os/exec"

	"github.com/openset/leetcode/internal/base"
)

var CmdPage = &base.Command{
	Run:       runPage,
	UsageLine: "page",
	Short:     "build index.md file",
	Long:      "build page file(index.md).",
}

func runPage(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	err := exec.Command(base.CmdName, "readme", "page").Run()
	base.CheckErr(err)
}
