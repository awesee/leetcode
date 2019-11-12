package problem167

func twoSum(numbers []int, target int) []int {
	s := 0
	l := len(numbers) - 1
	for s < l {
		r := numbers[s] + numbers[l]
		if r == target {
			return []int{s + 1, l + 1}
		} else if r < target {
			s++
		} else {
			l--
		}
	}
	return nil
}
