// Package clean implements the command clean.
package clean

import (
	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

// CmdClean - clean.CmdClean
var CmdClean = &base.Command{
	Run:       runClean,
	UsageLine: "clean",
	Short:     "remove cached files",
	Long:      "remove leetcode cached files.",
}

func runClean(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	leetcode.Clean()
}
