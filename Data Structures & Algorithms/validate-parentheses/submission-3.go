func isValid(s string) bool {
    stack := make([]byte, 0)

	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		if s[i] == ')' {
			if len(stack) == 0 {
					return false
				}
			if stack[len(stack) - 1] == '(' {
				
				stack = stack[ : len(stack) - 1]
			} else {
				return false
			}
		} else if s[i] == '}' {
			if len(stack) == 0 {
					return false
				}
			if stack[len(stack) - 1] == '{' {
				
				stack = stack[ : len(stack) - 1]
			} else {
				return false
			}
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack) - 1] == '[' {
				if len(stack) == 0 {
					return false
				}
				stack = stack[ : len(stack) - 1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
