class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        n = len(s1)
        m = len(s2)
        
        if n > m:
            return False
        check_cnt = {}
        for ch in s1:
            check_cnt[ch] = check_cnt.get(ch, 0) + 1
        
        cnt = {}
        for i in range(n):
            ch = s2[i]
            cnt[ch] = cnt.get(ch, 0) + 1
        
        if cnt == check_cnt:
            return True
        
        l = 0
        for r in range(n, m):
            right_ch = s2[r]
            cnt[right_ch] = cnt.get(right_ch, 0) + 1
            
            left_ch = s2[l]
            cnt[left_ch] -= 1
            if cnt[left_ch] == 0:
                del cnt[left_ch]  
            l += 1
            
            if cnt == check_cnt:
                return True
        
        return False