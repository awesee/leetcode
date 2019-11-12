package problem679

import "sort"

// 无法算出 24 的数字组合
var bad = map[int]bool{
	1111: true, 1112: true, 1113: true, 1114: true, 1115: true,
	1116: true, 1117: true, 1119: true, 1122: true, 1123: true,
	1124: true, 1125: true, 1133: true, 1159: true, 1167: true,
	1177: true, 1178: true, 1179: true, 1189: true, 1199: true,
	1222: true, 1223: true, 1299: true, 1355: true, 1499: true,
	1557: true, 1558: true, 1577: true, 1667: true, 1677: true,
	1678: true, 1777: true, 1778: true, 1899: true, 1999: true,
	2222: true, 2226: true, 2279: true, 2299: true, 2334: true,
	2555: true, 2556: true, 2599: true, 2677: true, 2777: true,
	2779: true, 2799: true, 2999: true, 3358: true, 3467: true,
	3488: true, 3555: true, 3577: true, 4459: true, 4466: true,
	4467: true, 4499: true, 4779: true, 4999: true, 5557: true,
	5558: true, 5569: true, 5579: true, 5777: true, 5778: true,
	5799: true, 5899: true, 5999: true, 6667: true, 6677: true,
	6678: true, 6699: true, 6777: true, 6778: true, 6779: true,
	6788: true, 6999: true, 7777: true, 7778: true, 7779: true,
	7788: true, 7789: true, 7799: true, 7888: true, 7899: true,
	7999: true, 8888: true, 8889: true, 8899: true, 8999: true,
	9999: true,
}

func judgePoint24(nums []int) bool {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < 4; i++ {
		sum = sum*10 + nums[i]
	}
	return !bad[sum]
}
