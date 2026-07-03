class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        cnt = {}
        l = 0
        r = 0
        mx_freq = 0
        mx = 0
        while r < len(s):
            cnt[s[r]] = cnt.get(s[r], 0) + 1
            mx_freq = max(cnt[s[r]], mx_freq)

            while (r - l + 1) - mx_freq  > k:
                    cnt[s[l]] -= 1
                    mx_freq = max(cnt.values())
                    l += 1                
            mx = max(mx, r - l + 1)
            r += 1
        return mx 

