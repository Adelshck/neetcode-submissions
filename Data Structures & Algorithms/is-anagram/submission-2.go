func isMapEqual(map1, map2 map[byte]int) bool {
	for i := range map1 {
		if map1[i] != map2[i] {
			return false
		}
	}
	return true
}
func isAnagram(s string, t string) bool {
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		map1[s[i]]++
		map2[t[i]]++
	}

	return isMapEqual(map1, map2)
}

