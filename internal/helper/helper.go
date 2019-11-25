// Package helper implements the command helper.
package helper

import (
	"bytes"
	"os/exec"

	"github.com/openset/leetcode/internal/base"
)

// CmdHelper - help.CmdHelper
var CmdHelper = &base.Command{
	Run:       runHelper,
	UsageLine: "helper",
	Short:     "build helper file",
	Long:      "build helper README.md file.",
	Hidden:    true,
}

func runHelper(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	var buf bytes.Buffer
	buf.WriteString(base.AuthInfo("helper"))
	buf.WriteString("\n# Helper\n\n")
	buf.WriteString("```text\n")
	out, err := exec.Command(base.CmdName).Output()
	base.CheckErr(err)
	buf.Write(out)
	buf.WriteString("```\n")
	base.FilePutContents("helper/README.md", buf.Bytes())
}
