package readme

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdReadme = &base.Command{
	Run:       runReadme,
	UsageLine: "readme",
	Short:     "build README.md file",
	Long:      "build README.md file.",
}

const (
	filename = "README.md"
	lineNum  = 12
)

func runReadme(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	wb := bytes.NewBuffer(nil)
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	base.CheckErr(err)
	defer f.Close()
	rb := bufio.NewReader(f)
	for ln := 0; ln < lineNum; ln++ {
		line, _, err := rb.ReadLine()
		if err == io.EOF {
			break
		}
		base.CheckErr(err)
		wb.Write(line)
		wb.WriteString("\n")
	}
	writeProblems(wb)
	err = ioutil.WriteFile(filename, wb.Bytes(), 0644)
	base.CheckErr(err)
}

func writeProblems(wb *bytes.Buffer) {
	problems := leetcode.ProblemsAll()
	problemsSet := make(map[int]string)
	maxId := 0
	for _, problem := range problems.StatStatusPairs {
		id := problem.Stat.FrontendQuestionId
		title := strings.TrimSpace(problem.Stat.QuestionTitle)
		slug := problem.Stat.QuestionTitleSlug
		levelName := problem.Difficulty.LevelName()
		format := "| %d | [%s](https://leetcode.com/problems/%s) | [Go](https://github.com/openset/leetcode/tree/master/solution/%s) | %s |\n"
		problemsSet[id] = fmt.Sprintf(format, id, title, slug, slug, levelName)
		if id > maxId {
			maxId = id
		}
	}
	for i := 0; i <= maxId; i++ {
		if row, ok := problemsSet[i]; ok {
			wb.WriteString(row)
		}
	}
}
