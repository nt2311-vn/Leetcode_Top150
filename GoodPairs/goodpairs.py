from typing import List


class Solution:
    def numIdenticalPairs(self, nums: List[int]) -> int:
        if len(nums) <= 1:
            return 0

        count = 0
        start = 0

        while start < len(nums):
            pointer = start + 1

            while pointer < len(nums):
                if nums[pointer] == nums[start]:
                    count += 1

                pointer += 1

            start += 1

        return count
