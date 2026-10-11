func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	ans := make([]int, n)
	stack := make([][2]int, 0, n)

	for i, v := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, [2]int{v, i})
			continue
		}
		if stack[len(stack) - 1][0] < v {
			for stack[len(stack) - 1][0] < v {
				el := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
				ans[el[1]] = i - el[1]
				if len(stack) == 0 {
					break
				}
			}
			stack = append(stack, [2]int{v, i})
			continue
		} else {
			stack = append(stack, [2]int{v, i})
		}
	}
	return ans
}
