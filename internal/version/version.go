package version

import (
	"fmt"
	"runtime"

	"github.com/openset/leetcode/internal/base"
)

const version = "0.0.1"

var CmdVersion = &base.Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     fmt.Sprintf("print %s version", base.CmdName),
	Long:      fmt.Sprintf("Version prints the %s version.", base.CmdName),
}

func runVersion(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Printf("%s version %s %s/%s\n", base.CmdName, version, runtime.GOOS, runtime.GOARCH)
}
