package base

import (
	"flag"
	"fmt"
)

// Run - base.Run
func Run() {
	flag.Usage = Usage
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	args := flag.Args()
	cmdName := flag.Arg(0)
	for _, cmd := range Commands {
		if cmd.Name() == cmdName {
			cmd.Run(cmd, args[1:])
			return
		}
	}
	fmt.Printf("%s %s: unknown command\n\n", CmdName, cmdName)
	fmt.Printf("Run '%s help' for usage.\n", CmdName)
}
