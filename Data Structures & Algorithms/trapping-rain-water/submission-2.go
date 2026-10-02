func trap(height []int) int {
	n := len(height)
	left := make([]int, n)
	curhl := 0
	right := make([]int, n)
	curhr := 0
	sum := 0
	for i:= 0; i < n; i++ {
		curhl = max(curhl, height[i])
		cur := curhl - height[i]
		if cur > 0 {
			left[i] = cur
		}
	}
	for i := n - 1; i > -1; i-- {
		curhr = max(curhr, height[i])
		cur := curhr - height[i]
		if cur > 0 {
			right[i] = cur
		}
	}

	for i := 0; i < n; i++ {
		sum += min(left[i], right[i])
	}
	return sum
}
