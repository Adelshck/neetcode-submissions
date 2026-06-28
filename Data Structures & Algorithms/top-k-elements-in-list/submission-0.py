class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        cnt = {}
        n = 0
        ans = []
        for i in nums:
            cnt[i] = cnt.get(i, 0) + 1
            n += 1
        bucket = [0] * (n + 1)
        for i, j in cnt.items():
            bucket[j] = i
        for i in range(n , -1 , -1):
            if bucket[i] == 0:
                continue
            ans.append(bucket[i])
            k -= 1
            if k == 0 :
                return ans
        
        