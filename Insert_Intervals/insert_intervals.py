from typing import List


class Solution:
    def insert(
        self, intervals: List[List[int]], newInterval: List[int]
    ) -> List[List[int]]:
        merge_intervals: List[List[int]] = []

        index = 0

        while index < len(intervals) and intervals[index][1] < newInterval[0]:
            merge_intervals.append(intervals[index])
            index += 1

        while index < len(intervals) and intervals[index][0] <= newInterval[1]:
            newInterval[0] = min(newInterval[0], intervals[index][0])
            newInterval[1] = max(newInterval[1], intervals[index][1])
            index += 1

        merge_intervals.append(newInterval)

        while index < len(intervals):
            merge_intervals.append(intervals[index])
            index += 1

        return merge_intervals
