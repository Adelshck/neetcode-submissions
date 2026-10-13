func largestRectangleArea(heights []int) int {
	stack := make([]int, 0, len(heights))
	mx := 0

	for i, v := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > v {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			mx = max(mx, heights[mid]*width)
		}

		stack = append(stack, i)
	}

	for len(stack) > 0 {
		mid := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		width := len(heights)
		if len(stack) > 0 {
			width = len(heights) - stack[len(stack)-1] - 1
		}

		mx = max(mx, heights[mid]*width)
	}

	return mx
}