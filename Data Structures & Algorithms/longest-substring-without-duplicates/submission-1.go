func lengthOfLongestSubstring(s string) int {
    seen := make(map[byte]bool)
    l := 0
    mx := 0

    for r := 0; r < len(s); r++ {
        for seen[s[r]] {
            delete(seen, s[l]) 
            l++
        }

        seen[s[r]] = true

        mx = max(mx, r-l+1)
    }

    return mx
}
