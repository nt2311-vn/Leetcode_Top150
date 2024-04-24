from typing import List


class Solution:
    def sortArrayByParityII(self, nums: List[int]) -> List[int]:
        sort_parity: List[int] = [0] * len(nums)

        i, j = 0, 1

        for num in nums:
            if num % 2 == 0:
                sort_parity[i] = num
                i += 2
            else:
                sort_parity[j] = num
                j += 2

        return sort_parity
