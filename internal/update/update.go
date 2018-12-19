package update

import (
	"os"
	"os/exec"

	"github.com/openset/leetcode/internal/base"
)

var CmdUpdate = &base.Command{
	Run:       runUpdate,
	UsageLine: "update",
	Short:     "update self",
	Long:      "automates testing the packages.",
}

func runUpdate(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	err := exec.Command("go", "get", "-u", "github.com/openset/leetcode").Run()
	base.CheckErr(err)
	c := exec.Command(base.CmdName, "version")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err = c.Run()
	base.CheckErr(err)
}
