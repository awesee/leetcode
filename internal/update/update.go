// Package update implements the command update.
package update

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/openset/leetcode/internal/base"
)

// CmdUpdate - update.CmdUpdate
var CmdUpdate = &base.Command{
	Run:       runUpdate,
	UsageLine: "update",
	Short:     "update or upgrade self",
	Long:      fmt.Sprintf("fetch the newest version of %s.", base.CmdName),
}

func runUpdate(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	err := exec.Command("go", "get", "-u", "github.com/openset/leetcode").Run()
	base.CheckErr(err)
	c := exec.Command(base.CmdName, "version")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Run()
	base.CheckErr(err)
}
