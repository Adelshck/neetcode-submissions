func twoSum(nums []int, target int) []int {
	targets := make(map[int]int, len(nums))
/*
a + b = targets 
b = targets - a

*/
	for i := range nums {
		diff := target - nums[i]
        if prevIdx, exist := targets[diff]; exist {
            return []int{prevIdx, i}
        }
        targets[nums[i]] = i
	}
    return []int{0, 0}
}
