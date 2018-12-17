package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/version"
)

func init() {
	base.Commands = []*base.Command{
		version.CmdVersion,
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", base.CmdName))

	flag.Usage = base.Usage
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
	}
	args := flag.Args()
	cmdName := args[0]
	for _, cmd := range base.Commands {
		if cmd.Name() == cmdName {
			cmd.Run(cmd, args[1:])
			base.Exit()
		}
	}
	fmt.Printf("%s %s: unknown command\n", base.CmdName, cmdName)
	fmt.Printf("Run '%s help' for usage.\n", base.CmdName)
}
