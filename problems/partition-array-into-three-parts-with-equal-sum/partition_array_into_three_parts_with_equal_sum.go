package problem1013

func canThreePartsEqualSum(A []int) bool {
	avg, sum := 0, 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	avg, sum, count := sum/3, 0, 0
	for _, v := range A {
		sum += v
		if sum == avg {
			count, sum = count+1, 0
		}
	}
	return count == 3
}
