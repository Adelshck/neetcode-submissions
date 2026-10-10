func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
		} else {
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]


			switch token {
			case "/":
				stack = append(stack, num2 / num1)
			case "*":
				stack = append(stack, num2 * num1)
			case "-":
				stack = append(stack, num2 - num1)
			case "+": 
				stack = append(stack, num2 + num1)
			}
		}
	}
	return stack[0]
}
