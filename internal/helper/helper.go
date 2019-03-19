package helper

import (
	"bytes"
	"os/exec"

	"github.com/openset/leetcode/internal/base"
)

var CmdHelper = &base.Command{
	Run:       runHelper,
	UsageLine: "helper",
	Short:     "build helper file",
	Long:      "build helper README.md file.",
}

func runHelper(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	var buf bytes.Buffer
	buf.WriteString(base.AuthInfo("helper"))
	buf.WriteString("\n# Helper\n\n")
	buf.WriteString("```text\n")
	out, _ := exec.Command(base.CmdName).Output()
	buf.Write(out)
	buf.WriteString("```\n")
	base.FilePutContents("helper/README.md", buf.Bytes())
}
