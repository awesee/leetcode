package problem989

func addToArrayForm(A []int, K int) []int {
	ans, carry, l := make([]int, 0), 0, len(A)-1
	for carry > 0 || K > 0 || l >= 0 {
		x := K%10 + carry
		if l >= 0 {
			x, l = x+A[l], l-1
		}
		carry, K = x/10, K/10
		ans = append([]int{x % 10}, ans...)
	}
	return ans
}
