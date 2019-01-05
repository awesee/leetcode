package kit

import (
	"reflect"
	"sort"
)

// IsEqualSliceInt returns a boolean
// reporting whether the slices values
// are deeply equal in any order.
func IsEqualSliceInt(s1, s2 []int) bool {
	sort.Ints(s1)
	sort.Ints(s2)
	return reflect.DeepEqual(s1, s2)
}
