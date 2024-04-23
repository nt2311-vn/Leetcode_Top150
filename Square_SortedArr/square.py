from typing import List


class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        sorted_square: List[int] = []
        i, j = 0, len(nums) - 1

        while i <= j:
            i_sqr = nums[i] ** 2
            j_sqr = nums[j] ** 2

            if i_sqr > j_sqr:
                sorted_square.append(i_sqr)
                i += 1
            else:
                sorted_square.append(j_sqr)
                j -= 1

        return sorted_square[::-1]
