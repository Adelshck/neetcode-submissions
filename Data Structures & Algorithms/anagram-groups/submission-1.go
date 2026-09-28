func countArrWord(str string) [26]int {
	var arrWord [26]int
	for i := range str {
		(arrWord[str[i]-'a'])++
	}
	return arrWord
}

func groupAnagrams(strs []string) [][]string {

	words := make(map[[26]int][]string)
	for _, v := range strs {
		arrWord := countArrWord(v)
		words[arrWord] = append(words[arrWord], v)
	}

	var ans [][]string
	for _, v := range words {
		ans = append(ans, v)
	}
	return ans
}