// Package page implements the command page.
package page

import (
	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/readme"
)

// CmdPage - page.CmdPage
var CmdPage = &base.Command{
	Run:       runPage,
	UsageLine: "page",
	Short:     "build index.md file",
	Long:      "build page file(index.md).",
	Hidden:    true,
}

func runPage(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	readme.CmdReadme.Run(readme.CmdReadme, []string{"page"})
}
