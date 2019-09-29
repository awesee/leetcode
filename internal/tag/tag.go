// Package tag implements the command tag.
package tag

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdTag = &base.Command{
	Run:       runTag,
	UsageLine: "tag",
	Short:     "build all tags file",
	Long:      "build all tags README.md file.",
}

func runTag(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	var wg sync.WaitGroup
	var buf bytes.Buffer
	buf.WriteString(base.AuthInfo("tag"))
	buf.WriteString("\n## 话题分类\n\n")
	buf.WriteString("| # | Title | 话题 | | # | Title | 话题 |\n")
	buf.WriteString("| :-: | - | :-: | - | :-: | - | :-: |\n")
	format := "| %d | [%s](https://github.com/openset/leetcode/tree/master/tag/%s/README.md) | [%s](https://openset.github.io/tags/%s/) | "
	n := buf.Len()
	for times := 0; times < 2; times++ {
		buf.Truncate(n)
		tags := leetcode.GetTags()
		for i, tag := range tags {
			if tag.Name != "" {
				fmt.Println(i+1, "\t"+tag.Name, "saving...")
			}
			buf.WriteString(fmt.Sprintf(format, i+1, tag.Name, tag.Slug, tag.TranslatedName, tag.Slug))
			if i&1 == 1 {
				buf.WriteString("\n")
			}
			wg.Add(1)
			go func(tag leetcode.TagType) {
				tag.SaveContents()
				wg.Done()
			}(tag)
		}
		if reflect.DeepEqual(tags, leetcode.GetTags()) {
			break
		}
	}
	base.FilePutContents("tag/README.md", buf.Bytes())
	wg.Wait()
}
