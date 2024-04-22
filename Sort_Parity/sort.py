from typing import List


class Solution:
    def sortArrayByParity(self, nums: List[int]) -> List[int]:

        if len(nums) == 1:
            return nums

        if len(nums) == 0:
            return []

        sort_parity: List[int] = [0] * len(nums)

        i, j = 0, len(nums) - 1

        for num in nums:
            if num % 2 == 0:
                sort_parity[i] = num
                i += 1
            else:
                sort_parity[j] = num
                j -= 1

        return sort_parity
