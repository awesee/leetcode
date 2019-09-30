// Package post implements the command post.
package post

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
		return
	}
	formatFilename := "%s-%s.md"
	formatTopicTag := "  [[%s](https://github.com/openset/leetcode/tree/master/tag/%s/README.md)]\n"
	formatSimilarQuestion := "  1. [%s](/problems/%s)%s\n"
	problems := leetcode.ProblemsAll()
	for _, problem := range problems.StatStatusPairs {
		questionId := problem.Stat.FrontendQuestionId
		titleSlug := problem.Stat.QuestionTitleSlug
		question := leetcode.QuestionData(titleSlug, false).Data.Question
		if question.TranslatedContent != "" {
			fmt.Println(questionId, "\t"+question.TranslatedTitle, "saving...")
			var buf bytes.Buffer
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
			buf.WriteString("{% raw %}\n\n")
			content := strings.ReplaceAll(question.TranslatedContent, "\r", "")
			// remove style
			reg := regexp.MustCompile(`<style[\S\s]+?</style>`)
			content = reg.ReplaceAllString(content, "")
			content = strings.ReplaceAll(content, "\n\n\t", "\n\t")
			content = strings.TrimSpace(content)
			buf.WriteString(content)
			buf.WriteString("\n\n{% endraw %}\n")
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
			buf.WriteString(fmt.Sprintf("\n## [解法](https://github.com/openset/leetcode/tree/master/problems/%s)\n", question.TitleSlug))
			filename := fmt.Sprintf(formatFilename, t.Format("2006-01-02"), question.TitleSlug)
			oldPath := filepath.Join(basePath, "leetcode", filename)
			newPath := filepath.Join(basePath, "_posts", filename)
			base.FilePutContents(oldPath, buf.Bytes())
			if leetcode.IsSolved(questionId) {
				_ = os.Rename(oldPath, newPath)
			} else {
				_ = os.Remove(newPath)
			}
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
		base.FilePutContents(filepath.Join(basePath, "_pages", filename), data)
	}
}

const frontMatter = `---
layout:     single
title:      "%s"
date:       %s +0800
categories: [Leetcode]
tags:       [%s]
permalink:  /problems/%s/
---
`

const tagTmpl = `---
title: "%s"
layout: tag
permalink: /tags/%s/
taxonomy: %s
---
`

var (
	homeDir, _ = os.UserHomeDir()
	basePath   = filepath.Join(homeDir, "openset", "openset")
)
