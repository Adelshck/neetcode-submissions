class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        numsSet = set(nums)
        mx = 0
        for i in numsSet:
            if i - 1 in numsSet:
                continue
            temp = 1
            while i + 1 in numsSet:
                temp += 1
                i += 1
            if temp > mx:
                mx = temp
        return mx
            
            
        