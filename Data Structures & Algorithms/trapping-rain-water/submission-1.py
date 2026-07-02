class Solution:
    def trap(self, height: List[int]) -> int:
        n = len(height)
        l = 0
        r = n - 1
        mxl = 0
        mxr = 0
        count = 0
        while l < r :
            if height[l] < height[r]:
                if height[l] >= mxl:
                    mxl = height[l]
                else:
                    count += mxl - height[l]
                l += 1
            else:
                if height[r] >= mxr:
                    mxr = height[r]
                else:
                    count += mxr - height[r]
                r -= 1
        return count