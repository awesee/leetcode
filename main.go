package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/description"
	"github.com/openset/leetcode/internal/help"
	"github.com/openset/leetcode/internal/helper"
	"github.com/openset/leetcode/internal/open"
	"github.com/openset/leetcode/internal/question"
	"github.com/openset/leetcode/internal/readme"
	"github.com/openset/leetcode/internal/test"
	"github.com/openset/leetcode/internal/update"
	"github.com/openset/leetcode/internal/version"
)

func init() {
	base.Commands = []*base.Command{
		readme.CmdReadme,
		question.CmdQuestion,
		open.CmdOpen,
		description.CmdDescription,
		test.CmdTest,
		update.CmdUpdate,
		helper.CmdHelper,
		version.CmdVersion,
		help.CmdHelp,
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
