package post

import (
	"bytes"
	"fmt"
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
		question := leetcode.QuestionData(titleSlug).Data.Question
		if question.TranslatedContent != "" {
			fmt.Println(question.QuestionFrontendId, "\t"+question.TranslatedTitle, "saving...")
			var buf bytes.Buffer
			questionId, _ := strconv.Atoi(question.QuestionFrontendId)
			t := time.Date(2016, 1, 1, 21, 30, 0, 0, time.Local)
			t = t.AddDate(0, 0, questionId)
			var tagSlugs []string
			var tagsBuf bytes.Buffer
			for _, tag := range question.TopicTags {
				tagSlugs = append(tagSlugs, tag.Slug)
				if tag.TranslatedName != "" {
					tag.Name = tag.TranslatedName
				}
				tagsBuf.WriteString(fmt.Sprintf(formatTopicTag, tag.Name, tag.Slug))
			}
			buf.WriteString(fmt.Sprintf(frontMatter,
				question.TranslatedTitle,
				t.Format("2006-01-02 15:04:05"),
				strings.Join(tagSlugs, ", "),
				question.TitleSlug,
			))
			buf.WriteString(fmt.Sprintf("\n## %s. %s%s\n\n", question.QuestionFrontendId, question.TranslatedTitle, question.Difficulty.Str()))
			content := strings.Replace(question.TranslatedContent, "\r", "", -1)
			// remove style
			reg := regexp.MustCompile(`<style[\S\s]+?</style>`)
			content = reg.ReplaceAllString(content, "")
			content = strings.Replace(content, "\n\n\t", "\n\t", -1)
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
			base.FilePutContents(path.Join("post", filename), buf.Bytes())
		}
	}
}

const frontMatter = `---
layout:     single
title:      "%s"
date:       %s +0800
categories: [leetcode]
tags:       [%s]
permalink:  /%s/
---
`
