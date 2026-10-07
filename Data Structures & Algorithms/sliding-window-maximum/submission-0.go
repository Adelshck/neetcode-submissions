func maxSlidingWindow(nums []int, k int) []int {
    var ans []int
	n := len(nums)
	l := 0
	idx_mx := 0
	for r := 0; r < k; r++ {
		if nums[idx_mx] < nums[r]{
			idx_mx = r
		}
	}

	ans = append(ans, nums[idx_mx])

	for r := k; r < n; r++ {
		l++
		if idx_mx < l {
			idx_mx = l
			for m := l; m <= r; m++ {
				if nums[idx_mx] < nums[m]{
					idx_mx = m
				}
			}
		} else {
			if nums[idx_mx] < nums[r]{
				idx_mx = r
			}
		}
		ans = append(ans, nums[idx_mx])
	}
	return ans
}
