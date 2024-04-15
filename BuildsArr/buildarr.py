from typing import List


class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:

        if len(nums) == 0:
            return []

        result_arr: List[int] = [0] * len(nums)

        i = 0

        while i < len(nums):
            result_arr[i] = nums[nums[i]]
            i += 1

        return result_arr
