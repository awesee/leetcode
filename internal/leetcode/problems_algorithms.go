package leetcode

import "github.com/openset/leetcode/internal/client"

func ProblemsAlgorithms() (ps problemsType) {
	data := remember(problemsAlgorithmsFile, 3, func() []byte {
		return client.Get(apiProblemsAlgorithmsUrl)
	})
	jsonDecode(data, &ps)
	return
}
