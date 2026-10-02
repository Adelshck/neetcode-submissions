import "slices"

func threeSum(nums []int) [][]int {
    n := len(nums)
    slices.Sort(nums)
    var ans [][]int

    for i := 0; i < n-2; i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        m := i + 1
        j := n - 1

        for m < j {
            sum := nums[i] + nums[m] + nums[j]

            if sum > 0 {
                j--
            } else if sum < 0 {
                m++
            } else {
                ans = append(ans, []int{nums[i], nums[m], nums[j]})
                m++
                j--

                for m < j && nums[m] == nums[m-1] {
                    m++
                }
                for m < j && nums[j] == nums[j+1] {
                    j--
                }
            }
        }
    }

    return ans
}