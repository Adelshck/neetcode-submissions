def anagram(s):
        arr = [0] * 26
        for i in s:
            arr[ord('a') - ord(i)] += 1
        return tuple(arr)
class Solution:

    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        cnt = defaultdict(list)
        ans = []
        for i in strs:
            cnt[anagram(i)].append(i)
        for i in cnt:
            ans.append(cnt[i])
        return ans
            
        