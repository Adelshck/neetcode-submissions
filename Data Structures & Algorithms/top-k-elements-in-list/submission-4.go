func topKFrequent(nums []int, k int) []int {
    n := len(nums)
    if k == n {
        return nums
    }
	ans := make([]int, 0, k)
	
	cnt := make(map[int]int, n)
    bucket := make([][]int, n + 1)
	for _, v := range nums {
		cnt[v]++
	}

	for val, cnt := range cnt {
        bucket[cnt] = append(bucket[cnt], val)
    }

    for k > 0 {
        if bucket[n] != nil {
            for _, v := range bucket[n] {
                ans = append(ans, v)
                k--
            }
            
        }
        n--
    }
    return ans
}
