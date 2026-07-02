class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        mx_size = 0
        n = len(s)
        cnt = set()
        l = 0
        r = 0
        while r < n:
            if s[r] not in cnt:
                cnt.add(s[r])
                r += 1
            else:
                while s[l] != s[r]:
                    cnt.remove(s[l])
                    l += 1
                cnt.remove(s[l])
                l += 1
            size = r - l
            mx_size = max(size, mx_size)
        return mx_size

