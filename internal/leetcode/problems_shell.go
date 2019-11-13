package leetcode

import "github.com/openset/leetcode/internal/client"

// ProblemsShell - leetcode.ProblemsShell
func ProblemsShell() (ps ProblemsType) {
	data := remember(problemsShellFile, 2, func() []byte {
		return client.Get(apiProblemsShellURL)
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
