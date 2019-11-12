package problem11

func maxArea(height []int) int {
	ans, l, r := 0, 0, len(height)-1
	for l < r {
		w, h := r-l, height[l]
		if h < height[r] {
			l++
		} else {
			h, r = height[r], r-1
		}
		if area := w * h; area > ans {
			ans = area
		}
	}
	return ans
}
