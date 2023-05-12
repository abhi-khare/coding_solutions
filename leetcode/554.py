from typing import List

class Solution:
    def leastBricks(self, wall: List[List[int]]) -> int:

        counter = []

        for idx, row in enumerate(wall):
            temp = []
            sum = 0
            for i in range(len(row)-1):
                sum += row[i]
                temp.append(sum)
            wall[idx] = temp
        
        for row in wall:
            for sep in row:
                counter.append(sep)
        
        print(counter)
        if len(counter) == 0:
            max_counter = 0
        else:
            max_counter = max(counter)

        
        return len(wall)-max_counter





wall = [[1,2,2,1],[3,1,2],[1,3,2],[2,4],[3,1,2],[1,3,1,1]]
sol = Solution()
ans = sol.leastBricks(wall)
print(ans)