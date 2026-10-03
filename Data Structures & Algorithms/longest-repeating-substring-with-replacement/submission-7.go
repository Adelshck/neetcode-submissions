func characterReplacement(s string, k int) int {
	seen := make(map[byte]int)
    mxSeen := 0
    mx, l := 0, 0

    for r := 0 ; r < len(s) ; r++ {
        seen[s[r]]++
        mxSeen = max(mxSeen, seen[s[r]])
        for (r - l + 1) - mxSeen > k {
            seen[s[l]]--
            l++
            mV := 0
            for _, v := range seen {
                mV = max(v, mV)
            }
            mxSeen = mV
        }
        mx = max(mx, (r - l + 1))
    }

    return mx
}
