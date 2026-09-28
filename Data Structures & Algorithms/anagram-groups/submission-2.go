func countArrWord(str string) [26]byte {
	var arrWord [26]byte
	for i := range str {
		(arrWord[str[i]-'a'])++
	}
	return arrWord
}

func groupAnagrams(strs []string) [][]string {

	words := make(map[[26]byte][]string)
	for _, v := range strs {
		arrWord := countArrWord(v)
		words[arrWord] = append(words[arrWord], v)
	}

	ans := make([][]string, 0, len(strs))
	for _, v := range words {
		ans = append(ans, v)
	}
	return ans
}