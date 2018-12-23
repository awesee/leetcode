package clean

import (
	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdClean = &base.Command{
	Run:       runClean,
	UsageLine: "clean",
	Short:     "clean cache file",
	Long:      "clean leetcode cache file.",
}

func runClean(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	leetcode.Clean()
}
