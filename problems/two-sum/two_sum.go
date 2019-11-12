package problem1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for j, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, j}
		}
		m[v] = j
	}
	return nil
}
