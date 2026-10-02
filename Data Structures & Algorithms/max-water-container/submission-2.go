func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1
	mx := 0
	for l < r {
		sq := (r - l) * min(heights[r], heights[l])
		mx = max(sq, mx)
		if heights[r] >= heights[l]{
			l++
		} else {
			r--
		}
	}
	return mx
}
