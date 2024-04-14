from typing import List


class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        summary_ranges: List[str] = []
        if len(nums) == 0:
            return summary_ranges

        index = 0

        while index < len(nums):
            start = nums[index]

            while index + 1 < len(nums) and nums[index + 1] == nums[index] + 1:
                index += 1

            if start == nums[index]:
                summary_ranges.append(start.__str__())
            else:
                summary_ranges.append(f"{start}->{nums[index]}")

            index += 1

        return summary_ranges
