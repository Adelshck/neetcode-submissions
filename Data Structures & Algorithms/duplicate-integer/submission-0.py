class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        counter = {}
        for i in nums:
            counter[i] = counter.get(i, 0) + 1
            if counter[i] == 2 :
                return True
        return False
        