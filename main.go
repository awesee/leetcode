// leetcode is a tool for managing leetcode source code.
package main

import (
	"github.com/awesee/leetcode/internal/base"
	"github.com/awesee/leetcode/internal/build"
	"github.com/awesee/leetcode/internal/clean"
	"github.com/awesee/leetcode/internal/description"
	"github.com/awesee/leetcode/internal/help"
	"github.com/awesee/leetcode/internal/helper"
	"github.com/awesee/leetcode/internal/open"
	"github.com/awesee/leetcode/internal/page"
	"github.com/awesee/leetcode/internal/post"
	"github.com/awesee/leetcode/internal/question"
	"github.com/awesee/leetcode/internal/readme"
	"github.com/awesee/leetcode/internal/tag"
	"github.com/awesee/leetcode/internal/test"
	"github.com/awesee/leetcode/internal/update"
	"github.com/awesee/leetcode/internal/version"
)

func main() {
	base.Register(
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
	)
	base.Run()
}
