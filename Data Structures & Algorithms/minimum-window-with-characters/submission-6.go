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
	idxs := make([][3]int, 0)
    cnt := make(map[byte]int)
	cnt_buf := make(map[byte]int)


	for i := range t {
		cnt[t[i]]++
	}

	for i := range cnt {
		if cnt[i] <= cnt_buf[i] {
			eqls++
		}
	}
	if eqls == m {
		return s
	}

	
	idxMin := 0;
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

			idx := [3]int{r - l + 2, r + 1, l - 1}
			idxs = append(idxs, idx)
		}
	}
	
	if len(idxs) == 0 {
    	return ""
	}
	for i := range idxs {
		if idxs[i][0] < 0 {
			continue
		}
		if idxs[idxMin][0] > idxs[i][0]{
			idxMin = i
		}
	}
	return s[idxs[idxMin][2]:idxs[idxMin][1]]
}