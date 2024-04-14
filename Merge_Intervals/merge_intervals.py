from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        merge_intervals: List[List[int]] = []

        if len(intervals) == 0:
            return merge_intervals

        intervals = sorted(intervals, key=lambda x: x[0])

        merge_intervals.append(intervals[0])

        index = 1

        while index < len(intervals):
            current = intervals[index]
            merge_bound = merge_intervals[len(merge_intervals) - 1]

            if current[0] <= merge_bound[1]:
                merge_bound[1] = max(merge_bound[1], current[1])

            else:
                merge_intervals.append(current)

            index += 1

        return merge_intervals
