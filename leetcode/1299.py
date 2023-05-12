from typing import List

class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        
        largest_ele = arr[-1]
        arr[-1] = -1
        for idx in range(len(arr)-2,-1,-1):
            
            temp = arr[idx]

            arr[idx] = largest_ele
            if temp > largest_ele:
                largest_ele = temp
            

        return arr


nums = [17,15]

sol = Solution()
out = sol.replaceElements(nums)
print(out)

