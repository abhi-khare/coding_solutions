from typing import List
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        if numRows == 1:
            return [[1]]
        elif numRows == 2:
            return [[1], [1,1]]
        else:
            ans = self.generate(numRows-1)
            temp = [1]
            for i in range(0,len(ans[-1])-1):
                temp.append(ans[-1][i] + ans[-1][i+1])
            temp.append(1)
            ans.append(temp)
            return ans




n = 5
sol = Solution()
ans = sol.generate(n)
print(ans)