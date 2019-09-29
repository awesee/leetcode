// Package version implements the command version.
package version

import (
	"fmt"
	"runtime"

	"github.com/openset/leetcode/internal/base"
)

const version = "1.4.8"

var CmdVersion = &base.Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     fmt.Sprintf("print %s version", base.CmdName),
	Long:      fmt.Sprintf("prints the %s version.", base.CmdName),
}

func runVersion(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	fmt.Printf("%s version %s %s/%s\n", base.CmdName, version, runtime.GOOS, runtime.GOARCH)
}

func String() string {
	return version
}
