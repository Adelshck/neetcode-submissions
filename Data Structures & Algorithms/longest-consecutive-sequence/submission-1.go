func longestConsecutive(nums []int) int {
	nummap := make(map[int]bool)
    mx := 0
	for _, v := range nums {
		nummap[v] = true
	}

    for i := range nummap {
        if _, exist := nummap[i - 1]; exist {
            continue
        }
        temp := 1

        if _, exist := nummap[i + 1]; exist {
            temp++
            currentNum := i + 1
            for {
                _, stillExists := nummap[currentNum+1]
                if !stillExists {
                    break 
                }
                currentNum++ 
                temp++
            }
            
        }
        mx = max(temp, mx)
    }
    return mx
}
