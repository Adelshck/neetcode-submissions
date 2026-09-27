func isAnagram(s string, t string) bool {
	check := make(map[byte]int)
	for i := range s {
		check[s[i]]++
	}
	for i := range t {
		check[t[i]]--
	}
	for i,_ := range check{
		if check[i] != 0 {
			return false
		}
	}
	return true
}
