class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [1] * n
        k = 1
        for i in range(0, n - 1):
            k *= nums[i]
            ans[i + 1] *= k
        k = 1
        for i in range(n - 1, 0 , -1):
            k *= nums[i]
            ans[i - 1] *= k
        return ans