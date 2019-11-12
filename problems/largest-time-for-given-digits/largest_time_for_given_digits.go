package problem949

import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(A []int) string {
	sort.Ints(A)
	for i := 3; i >= 0; i-- {
		if A[i] < 3 {
			for j := 3; j >= 0; j-- {
				if j != i && A[i]*10+A[j] < 24 {
					for k := 3; k >= 0; k-- {
						if k != i && k != j && A[k] < 6 {
							l := 6 - i - j - k
							return fmt.Sprintf("%d%d:%d%d", A[i], A[j], A[k], A[l])
						}
					}
				}
			}
		}
	}
	return ""
}
