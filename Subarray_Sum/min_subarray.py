from typing import List


class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        min_len = len(nums) + 1
        sum = 0
        start = 0
        end = 0

        for num in nums:
            sum += num

            while sum >= target:
                if min_len > end - start + 1:
                    min_len = end - start + 1

                sum -= nums[start]
                start += 1

            end += 1

        if min_len == len(nums) + 1:
            return 0

        return min_len
