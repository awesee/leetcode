package leetcode

import "github.com/awesee/leetcode/internal/client"

// ProblemsAlgorithms - leetcode.ProblemsAlgorithms
func ProblemsAlgorithms() (ps ProblemsType) {
	data := remember(problemsAlgorithmsFile, 2, func() []byte {
		return client.Get(apiProblemsAlgorithmsURL)
	})
	jsonDecode(data, &ps)
	return
}
