from typing import List


class Solution:
    def arithmeticTriplets(self, nums: List[int], diff: int) -> int:
        if len(nums) < 3:
            return 0

        triplets = 0
        for num in nums[: len(nums) - 2]:
            equal_diff, compare_num = 0, num
            for next_num in nums[1:]:
                if next_num - compare_num == diff:
                    equal_diff += 1
                    compare_num = next_num

                if equal_diff == 2:
                    triplets += 1
                    break

        return triplets
