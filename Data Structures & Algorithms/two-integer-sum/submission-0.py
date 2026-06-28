class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        n = len(nums)
        cnt = {}
        for i in range(n):
            
            diff = target - nums[i]
            if cnt.get(diff, -1) != -1:
                return [cnt[diff], i]  
            cnt[nums[i]] = i    
