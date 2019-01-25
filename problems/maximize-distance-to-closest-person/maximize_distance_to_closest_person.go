package maximize_distance_to_closest_person

func maxDistToClosest(seats []int) int {
	ans, prev, future, l := 0, -1, 0, len(seats)
	for i, v := range seats {
		if v == 1 {
			prev = i
		} else {
			for future < l && seats[future] == 0 || future < i {
				future++
			}
			left, right := l, l
			if prev != -1 {
				left = i - prev
			}
			if future != l {
				right = future - i
			}
			if right < left {
				left = right
			}
			if left > ans {
				ans = left
			}
		}
	}
	return ans
}
