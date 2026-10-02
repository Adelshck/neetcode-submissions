func trap(height []int) int {
    l, r := 0, len(height)-1
    maxL, maxR := 0, 0
    ans := 0

    for l < r {
        if height[l] < height[r] {
            if height[l] >= maxL {
                maxL = height[l]
            } else {
                ans += maxL - height[l]
            }
            l++
        } else {
            if height[r] >= maxR {
                maxR = height[r]
            } else {
                ans += maxR - height[r]
            }
            r--
        }
    }
    return ans
}