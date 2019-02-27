package leetcode

import "github.com/openset/leetcode/internal/client"

func ProblemsShell() (ps problemsType) {
	data := remember(problemsShellFile, 2, func() []byte {
		return client.Get(apiProblemsShellUrl)
	})
	jsonDecode(data, &ps)
	return
}

func init() {
	problems := ProblemsShell()
	for _, problem := range problems.StatStatusPairs {
		saveAsBash(problem.Stat.QuestionTitleSlug)
	}
}
