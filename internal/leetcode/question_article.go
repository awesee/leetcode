package leetcode

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/openset/leetcode/internal/client"
)

func GetDescription(articleSlug string) string {
	fmt.Println("\tquestion article", "saving...")
	filename := fmt.Sprintf(questionArticleFile, articleSlug)
	html := remember(filename, 30, func() []byte {
		return client.Get(fmt.Sprintf(questionArticleUrl, articleSlug))
	})
	reg := regexp.MustCompile(`<div class="block-markdown question">([\S\s]+?)</div>`)
	matches := reg.FindSubmatch(html)
	if len(matches) >= 2 {
		return string(bytes.TrimSpace(matches[1]))
	}
	return ""
}
