func isValid(s string) bool {
    stack := make([]byte, 0)

    pairs := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for i := range s {
        if s[i] == '(' || s[i] == '{' || s[i] == '[' {
            stack = append(stack, s[i])
            continue
        }

        if len(stack) == 0 {
            return false
        }

        if stack[len(stack)-1] != pairs[s[i]] {
            return false
        }

        stack = stack[:len(stack)-1]
    }

    return len(stack) == 0
}