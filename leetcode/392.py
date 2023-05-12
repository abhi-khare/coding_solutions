class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:

        if len(t)< len(s):
            return False
        
        idx, jdx = 0,0

        while idx < len(s) and jdx< len(t):
            if s[idx] == t[jdx]:
                idx+=1
            jdx+=1
            
        
        if idx != len(s):
            return False
        else:
            return True




s,t = "bb", "b"

sol = Solution()
out = sol.isSubsequence(s,t)

print(out)