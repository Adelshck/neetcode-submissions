func productExceptSelf(nums []int) []int {
	n := len(nums)
	pref, suff := 1, 1
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}
	for i, v := range nums {
		ans[i] *= pref
		ans[n - i - 1] *= suff
		pref *= v
		suff *= nums[n - i - 1]
	}
	return ans
}
