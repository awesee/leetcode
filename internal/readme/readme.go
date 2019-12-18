// Package readme implements the command readme.
package readme

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

const defaultStr = `
# [LeetCode](https://openset.github.io/leetcode)
LeetCode Problems' Solutions
[[力扣](https://openset.github.io/categories/leetcode/) ∙ [话题分类](https://github.com/openset/leetcode/blob/master/tag/README.md)]

[![Build Status](https://travis-ci.org/openset/leetcode.svg?branch=master)](https://travis-ci.org/openset/leetcode)
[![codecov](https://codecov.io/gh/openset/leetcode/branch/master/graph/badge.svg)](https://codecov.io/gh/openset/leetcode)
[![Go Report Card](https://goreportcard.com/badge/github.com/openset/leetcode)](https://goreportcard.com/report/github.com/openset/leetcode)
[![GitHub contributors](https://img.shields.io/github/contributors/openset/leetcode.svg)](https://github.com/openset/leetcode/graphs/contributors)
[![license](https://img.shields.io/github/license/openset/leetcode.svg)](https://github.com/openset/leetcode/blob/master/LICENSE)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fopenset%2Fleetcode.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fopenset%2Fleetcode?ref=badge_shield)
[![Join the chat](https://badges.gitter.im/openset/leetcode.svg)](https://gitter.im/openset/leetcode?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

`

// CmdReadme - readme.CmdReadme
var CmdReadme = &base.Command{
	Run:       runReadme,
	UsageLine: "readme",
	Short:     "build README.md file",
	Long:      "build README.md file.",
	Hidden:    true,
}

var (
	buildCmd = "readme"
	fileName = "README.md"
	maxID    = 0
	pageSize = 300
	step     = 50
	num      = 6
)

func runReadme(cmd *base.Command, args []string) {
	if len(args) == 1 && args[0] == "page" {
		buildCmd = "page"
		fileName = "index.md"
		leetcode.LockStr = " &hearts;"
	}
	if len(args) > 1 {
		cmd.Usage()
		return
	}
	var buf bytes.Buffer
	buf.WriteString(base.AuthInfo(buildCmd))
	buf.WriteString(defaultStr)
	writeProblems(&buf)
	base.FilePutContents(fileName, buf.Bytes())
}

func writeProblems(buf *bytes.Buffer) {
	problems := leetcode.ProblemsAll().StatStatusPairs
	count := len(problems)
	if count > 0 {
		maxID = problems[0].Stat.FrontendQuestionID
		writeNav(buf)
		// list
		buf.WriteString("| # | Title | Solution | Difficulty |\n")
		buf.WriteString("| :-: | - | - | :-: |\n")
		n := buf.Len()
		for i := 1; i < maxID/pageSize; i++ {
			for problems[count-1].Stat.FrontendQuestionID <= pageSize*i {
				count--
				problems[count].WriteRow(buf, "../problems")
			}
			fileName := filepath.Join("readme", fmt.Sprintf("%d-%d.md", pageSize*(i-1)+1, pageSize*i))
			base.FilePutContents(fileName, buf.Bytes())
			buf.Truncate(n)
		}
		for _, problem := range problems[0:count] {
			problem.WriteRow(buf, "problems")
		}
	}
}

func writeNav(buf *bytes.Buffer) {
	// table
	buf.WriteString("<table><thead>\n")
	for i := 0; i < maxID; i += step * num {
		buf.WriteString("<tr>\n")
		for j := 0; j < num; j++ {
			buf.WriteString(fmt.Sprintf("\t<th align=\"center\"><a href=\"%s\">[%d-%d]</a></th>\n", linkStr(i+j*step+step), 1+i+j*step, i+j*step+step))
		}
		buf.WriteString("</tr>\n")
	}
	buf.WriteString("</thead></table>\n\n")
}

func linkStr(num int) string {
	if num > maxID-maxID%pageSize-pageSize {
		return fmt.Sprintf("%s/README.md#%d", base.URL, num)
	}
	return fmt.Sprintf("%s/readme/%d-%d.md#%d", base.URL, (num-1)/pageSize*pageSize+1, ((num-1)/pageSize+1)*pageSize, num-step+1)
}
