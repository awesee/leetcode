package leetcode

import "github.com/openset/leetcode/internal/client"

func ProblemsDatabase() (ps problemsType) {
	data := remember(problemsDatabaseFile, 3, func() []byte {
		return client.Get(apiProblemsDatabaseUrl)
	})
	jsonDecode(data, &ps)
	return
}

func init() {
	problems := ProblemsDatabase()
	for _, problem := range problems.StatStatusPairs {
		slug := problem.Stat.QuestionTitleSlug
		langSet[slug] = "MySQL"
	}
}
