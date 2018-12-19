package help

import (
	"fmt"

	"github.com/openset/leetcode/internal/base"
)

var CmdHelp = &base.Command{
	Run:       runHelp,
	UsageLine: "help [command]",
	Short:     "show command usage",
	Long:      "show command usage detail.",
}

func runHelp(cmd *base.Command, args []string) {
	if len(args) < 1 {
		base.Usage()
	}
	cmdName := args[0]
	for _, cmd := range base.Commands {
		if cmd.Name() == cmdName {
			cmd.UsageHelp()
		}
	}
	fmt.Printf("%s help %s: unknown help topic.\n", base.CmdName, cmdName)
	fmt.Printf("Run '%s help' for usage.\n", base.CmdName)
}
