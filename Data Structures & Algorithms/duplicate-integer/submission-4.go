func hasDuplicate(nums []int) bool {
	ma := make(map[int]int)

	for _, v := range nums {
		ma[v]++
		if ma[v] == 2 {
			return true
		}
	}
	return false
}
