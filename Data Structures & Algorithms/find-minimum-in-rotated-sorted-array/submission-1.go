func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	res := nums[r]
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] <= nums[r] {
			r = mid - 1
		} else if nums[mid] > nums[r] {
			l = mid + 1
		}
		res = min(res, nums[mid])
	}
	return res
}