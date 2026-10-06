func minWindow(s string, t string) string {
	n, m := len(s), len(t)	
	if n == 1 && m == 1 {
		if s[0] == t[0] {
			return s
		} else {
			return ""
		}
	}
	eqls := 0
    cnt := make(map[byte]int)
	cnt_buf := make(map[byte]int)
	mxR, mxL, mxLen := 0, 0, n + 1

	for i := range t {
		cnt[t[i]]++
	}

	
	if n < m {
		return ""
	}

	l := 0

	for r := 0; r < n; r++ {
    	cnt_buf[s[r]]++

    	if cnt_buf[s[r]] <= cnt[s[r]] {
        	eqls++
    	}

    	if eqls == m {
        	for eqls == m {
				cnt_buf[s[l]]--

				if cnt_buf[s[l]] < cnt[s[l]] {
					eqls--
				}

				l++
			}

			if mxLen > (r - l + 2){
				mxR = r
				mxL = l - 1
				mxLen = r - l + 2
			}
		}
	}

	
	if mxLen == n+1 {
    	return ""
	}

	return s[mxL : mxR+1]
}