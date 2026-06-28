class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        k = len(s)
        if k != len(t):
            return False

        arr = [0] * 26
        for i in s:
            arr[ord('a') - ord(i)] += 1
        for i in t:
            arr[ord('a') - ord(i)] -= 1
        for i in arr:
            if i != 0:
                return False
        return True

    
        