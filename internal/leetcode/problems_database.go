package leetcode

import "github.com/openset/leetcode/internal/client"

// ProblemsDatabase - leetcode.ProblemsDatabase
func ProblemsDatabase() (ps ProblemsType) {
	data := remember(problemsDatabaseFile, 2, func() []byte {
		return client.Get(apiProblemsDatabaseURL)
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
