package problem1018

func prefixesDivBy5(A []int) []bool {
	ans, v := make([]bool, len(A)), 0
	for i, b := range A {
		v = (v<<1 + b) % 5
		ans[i] = v == 0
	}
	return ans
}
