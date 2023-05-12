from typing import List

class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:

        for n in nums:
            idx = abs(n)-1
            if nums[idx]>0:
                nums[idx] = -1 * nums[idx]
        
        res = []
        for idx, n in enumerate(nums):
            if n >0:
                res.append(idx+1)
        return res



nums = [1]

sol =Solution()
ans = sol.findDisappearedNumbers(nums)
print(ans)