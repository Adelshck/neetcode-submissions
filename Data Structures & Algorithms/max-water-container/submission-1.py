def area(l, r, a):
    return (r - l) * min(a[l], a[r])
class Solution:
    def maxArea(self, heights: List[int]) -> int:
        n = len(heights)
        l = 0
        r = n - 1
        mx = -1
        while(l < r):
            if area(l , r, heights) > mx:
                mx = area(l , r, heights)
            if heights[l] <= heights[r]:
                l += 1
            elif heights[l] > heights[r]:
                r -= 1
        return mx
        