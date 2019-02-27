package leetcode

import "github.com/openset/leetcode/internal/client"

func ProblemsDatabase() (ps problemsType) {
	data := remember(problemsDatabaseFile, 2, func() []byte {
		return client.Get(apiProblemsDatabaseUrl)
	})
	jsonDecode(data, &ps)
	return
}

func init() {
	problems := ProblemsDatabase()
	for _, problem := range problems.StatStatusPairs {
		saveAsMySQL(problem.Stat.QuestionTitleSlug)
	}
}
