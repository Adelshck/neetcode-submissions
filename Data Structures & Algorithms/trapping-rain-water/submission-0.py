class Solution:
    def trap(self, height: List[int]) -> int:
        count = 0
        arr = []
        n = len(height)
        left_max = [0] * n
        right_max = [0] * n
        mxl = 0
        mxr = 0
        for i in range(n):
            if mxl < height[i]:
                mxl = height[i]
            left_max[i] = mxl
        for i in range(n - 1, -1, -1):
            if mxr < height[i]:
                mxr = height[i]
            right_max[i] = mxr
        for i in range(n):
            high = min(right_max[i], left_max[i])
            count += high - height[i]
        return count