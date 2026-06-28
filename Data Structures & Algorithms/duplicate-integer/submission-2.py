class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        counter = set()
        for i in nums:
            if i in counter:
                return True
            counter.add(i)
        return False
            
        