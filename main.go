// leetcode is a tool for managing leetcode source code.
package main

import (
	"flag"
	"fmt"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/build"
	"github.com/openset/leetcode/internal/clean"
	"github.com/openset/leetcode/internal/description"
	"github.com/openset/leetcode/internal/help"
	"github.com/openset/leetcode/internal/helper"
	"github.com/openset/leetcode/internal/open"
	"github.com/openset/leetcode/internal/page"
	"github.com/openset/leetcode/internal/post"
	"github.com/openset/leetcode/internal/question"
	"github.com/openset/leetcode/internal/readme"
	"github.com/openset/leetcode/internal/tag"
	"github.com/openset/leetcode/internal/test"
	"github.com/openset/leetcode/internal/update"
	"github.com/openset/leetcode/internal/version"
)

func init() {
	base.Commands = []*base.Command{
		readme.CmdReadme,
		page.CmdPage,
		tag.CmdTag,
		helper.CmdHelper,
		question.CmdQuestion,
		open.CmdOpen,
		test.CmdTest,
		description.CmdDescription,
		update.CmdUpdate,
		build.CmdBuild,
		clean.CmdClean,
		version.CmdVersion,
		help.CmdHelp,
		post.CmdPost,
	}
}

func main() {
	flag.Usage = base.Usage
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}
	args := flag.Args()
	cmdName := flag.Arg(0)
	for _, cmd := range base.Commands {
		if cmd.Name() == cmdName {
			cmd.Run(cmd, args[1:])
			return
		}
	}
	fmt.Printf("%s %s: unknown command\n\n", base.CmdName, cmdName)
	fmt.Printf("Run '%s help' for usage.\n", base.CmdName)
}
