from typing import List, Tuple


class Solution:
    def maximumCount(self, nums: List[int]) -> int:
        negative, postive = 0, 0
        lower, upper = binary_search(nums)

        index = 0

        while index < lower:
            if nums[index] < 0:
                negative += 1
            index += 1

        index = upper

        while index < len(nums):
            if index >= 0 and nums[index] > 0:
                postive += 1
            index += 1

        return max(negative, postive)


def binary_search(nums: List[int]) -> Tuple[int, int]:
    lower, upper = 0, len(nums) - 1

    while lower <= upper:
        mid = lower + (upper - lower) // 2

        if nums[mid] < 0:
            lower = mid + 1
        else:
            upper = mid - 1

    return lower, upper
