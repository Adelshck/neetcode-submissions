class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        k = len(s)
        if k != len(t):
            return False

        s_cnt = {}
        t_cnt = {}
        for i in range(k):
            s_cnt[s[i]] = s_cnt.get(s[i], 0) + 1
            t_cnt[t[i]] = t_cnt.get(t[i], 0) + 1

        return s_cnt == t_cnt
    
        