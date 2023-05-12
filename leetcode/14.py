from typing import List

class Solution:

    def longestCommonPrefix(self, strs: List[str]) -> str:

        strs.sort()

        str_a, str_b = strs[0], strs[-1]
        res = ""

        for idx in range(len(str_a)):

            if str_a[idx]==str_b[idx]:
                res += str_a[idx]
            else:
                break
        
        return res




strs = ["dog","dogcar", "d"]

sol = Solution()

#ans = sol.common_prefix("a", "abhinash")
ans = sol.longestCommonPrefix(strs)
print(ans)