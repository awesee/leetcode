package post

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/leetcode"
)

var CmdPost = &base.Command{
	Run:       runPost,
	UsageLine: "post",
	Short:     "build post files",
	Long:      "build all post files.",
	Hidden:    true,
}

func runPost(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}
	formatFilename := "%s-%s.md"
	formatTopicTag := "  [[%s](https://github.com/openset/leetcode/tree/master/tag/%s/README.md)]\n"
	formatSimilarQuestion := "  1. [%s](/%s)%s\n"
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		titleSlug := problem.Stat.QuestionTitleSlug
		question := leetcode.QuestionData(titleSlug, false).Data.Question
		if question.TranslatedContent != "" {
			fmt.Println(question.QuestionFrontendId, "\t"+question.TranslatedTitle, "saving...")
			var buf bytes.Buffer
			questionId, _ := strconv.Atoi(question.QuestionFrontendId)
			t := time.Date(2016, 1, 1, 21, 30, 0, 0, time.Local)
			t = t.AddDate(0, 0, questionId)
			var tags []string
			var tagsBuf bytes.Buffer
			for _, tag := range question.TopicTags {
				tags = append(tags, tag.Name)
				if tag.TranslatedName != "" {
					tag.Name = tag.TranslatedName
				}
				tagsBuf.WriteString(fmt.Sprintf(formatTopicTag, tag.Name, tag.Slug))
			}
			buf.WriteString(fmt.Sprintf(frontMatter,
				question.TranslatedTitle,
				t.Format("2006-01-02 15:04:05"),
				strings.Join(tags, ", "),
				question.TitleSlug,
			))
			buf.WriteString(fmt.Sprintf("\n## %s. %s%s\n\n", question.QuestionFrontendId, question.TranslatedTitle, question.Difficulty.Str()))
			content := strings.ReplaceAll(question.TranslatedContent, "\r", "")
			// remove style
			reg := regexp.MustCompile(`<style[\S\s]+?</style>`)
			content = reg.ReplaceAllString(content, "")
			content = strings.ReplaceAll(content, "\n\n\t", "\n\t")
			content = strings.TrimSpace(content)
			buf.WriteString(content + "\n")
			if len(question.TopicTags) > 0 {
				buf.WriteString("\n### 相关话题\n")
			}
			buf.Write(tagsBuf.Bytes())
			sq := question.GetSimilarQuestion()
			if len(sq) > 0 {
				buf.WriteString("\n### 相似题目\n")
			}
			for _, q := range sq {
				if q.TranslatedTitle != "" {
					q.Title = q.TranslatedTitle
				}
				buf.WriteString(fmt.Sprintf(formatSimilarQuestion, q.Title, q.TitleSlug, q.Difficulty.Str()))
			}
			buf.WriteString("\n---\n")
			buf.WriteString(fmt.Sprintf("\n## [答案](https://github.com/openset/leetcode/tree/master/problems/%s)\n", question.TitleSlug))
			filename := fmt.Sprintf(formatFilename, t.Format("2006-01-02"), question.TitleSlug)
			postType := "leetcode"
			if inPosts[questionId] {
				postType = "_posts"
			}
			base.FilePutContents(path.Join(basePath, postType, filename), buf.Bytes())
		}
	}
	postTags()
}

func postTags() {
	for _, tag := range leetcode.GetTags() {
		title := tag.TranslatedName
		if title == "" {
			title = tag.Name
		}
		filename := fmt.Sprintf("tag-%s.md", tag.Slug)
		data := []byte(fmt.Sprintf(tagTmpl, title, tag.Slug, tag.Name))
		base.FilePutContents(path.Join(basePath, "_pages", filename), data)
	}
}

const frontMatter = `---
layout:     single
title:      "%s"
date:       %s +0800
categories: [Leetcode]
tags:       [%s]
permalink:  /%s/
---
`

const tagTmpl = `---
title: "%s"
layout: tag
permalink: /tags/%s/
taxonomy: %s
---
`

var homeDir, _ = os.UserHomeDir()

var basePath = path.Join(homeDir, "openset", "openset")

var inPosts = map[int]bool{
	1:    true,
	2:    true,
	3:    true,
	4:    true,
	7:    true,
	8:    true,
	9:    true,
	12:   true,
	13:   true,
	15:   true,
	18:   true,
	20:   true,
	58:   true,
	101:  true,
	139:  true,
	168:  true,
	171:  true,
	172:  true,
	190:  true,
	191:  true,
	198:  true,
	233:  true,
	746:  true,
	793:  true,
	849:  true,
	941:  true,
	984:  true,
	1000: true,
}
