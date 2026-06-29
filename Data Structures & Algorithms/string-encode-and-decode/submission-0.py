class Solution:

    def encode(self, strs: List[str]) -> str:
        word = ""
        for i in strs:
            n = len(i)
            word += f"{n}#{i}"
        return word            

    def decode(self, s: str) -> List[str]:
        ans = []
        i = 0
        while i < len(s):
            k = s.find('#', i)
            if k == -1:
                ans.append(s[i:])
                return ans
            num = int(s[i:k])
            ans.append(s[k + 1 : k + 1 + num])
            i = k + 1 + num
        return ans

