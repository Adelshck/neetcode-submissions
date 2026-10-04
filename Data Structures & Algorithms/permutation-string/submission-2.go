func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)

	var arrS1 [26]int
	var arrS2 [26]int

	if n > m {
		return false
	}

	for i := range s1 {
		arrS1[s1[i] - 'a']++
		arrS2[s2[i] - 'a']++
	}
	if arrS1 == arrS2 {
		return true
	}

	l := 0
	for r := n; r < m ; r++ {
		arrS2[s2[r] - 'a']++
		arrS2[s2[l] - 'a']--
		l++
		if arrS1 == arrS2 {
			return true
		}
	}
	return false
}
