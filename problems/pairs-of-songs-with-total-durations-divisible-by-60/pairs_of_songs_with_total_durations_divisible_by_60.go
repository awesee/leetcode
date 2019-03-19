package pairs_of_songs_with_total_durations_divisible_by_60

func numPairsDivisibleBy60(time []int) int {
	ans, m := 0, [60]int{}
	for _, t := range time {
		ans += m[(60-t%60)%60]
		m[t%60]++
	}
	return ans
}
