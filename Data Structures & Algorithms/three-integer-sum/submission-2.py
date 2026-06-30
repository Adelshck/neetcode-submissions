class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        n = len(nums)
        ans = []
        for i in range(n):
            l = 0
            r = n - 1
            while l < r:
                if nums[l] + nums[r] > -nums[i]:
                    r -= 1
                elif nums[l] + nums[r] < -nums[i]:
                    l += 1
                else:
                    if l != i and r != i:
                        a = sorted([nums[i], nums[l], nums[r]])
                        if a not in ans:
                            ans.append(a)
                    r -= 1
                    l += 1
        return ans
            
        