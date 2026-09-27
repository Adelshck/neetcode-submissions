func twoSum(nums []int, target int) []int {
	targets := make(map[int]int)
/*
a + b = targets 
b = targets - a

*/
	for i := range nums {
		diff := target - nums[i]
        if _, exist := targets[diff]; exist {
            return []int{targets[diff], i}
        }
        targets[nums[i]] = i
	}
    return []int{0, 0}
}
