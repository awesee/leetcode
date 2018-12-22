package base

import (
	"reflect"
	"sort"
)

func IsEqualSliceInt(s1, s2 []int) bool {
	sort.Ints(s1)
	sort.Ints(s2)
	return reflect.DeepEqual(s1, s2)
}
