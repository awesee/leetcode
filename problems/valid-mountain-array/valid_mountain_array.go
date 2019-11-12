package problem941

func validMountainArray(A []int) bool {
	flag := false
	if len(A) < 3 {
		return false
	}
	for i, v := range A[1:] {
		if !flag && v <= A[i] {
			if flag = true; i == 0 {
				return false
			}
		}
		if flag && v >= A[i] {
			return false
		}
	}
	return flag
}
