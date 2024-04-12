from typing import Dict, List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums) <= 1:
            return len(nums)

        mapNum: Dict[int, bool] = {}

        for num in nums:
            mapNum[num] = True

        longest_steak = 0

        for num in mapNum.keys():
            if num - 1 not in mapNum:
                current_streak = 1
                next_num = num + 1

                while next_num in mapNum:
                    current_streak += 1
                    next_num += 1

                if current_streak > longest_steak:
                    longest_steak = current_streak

        return longest_steak
