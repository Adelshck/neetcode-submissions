func isAlnum(b byte) bool {
    return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    l, r := 0, len(s)-1

    for l < r {
        for l < r && !isAlnum(s[l]) {
            l++
        }
        for l < r && !isAlnum(s[r]) {
            r--
        }

        if s[l] != s[r] {
            return false
        }
        l++
        r--
    }

    return true
}
