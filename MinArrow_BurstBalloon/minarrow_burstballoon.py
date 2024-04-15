from typing import List


class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        if len(points) <= 1:
            return len(points)

        points = sorted(points, key=lambda point: point[0])

        count = 1
        end = points[0][1]
        index = 1

        while index < len(points):
            if points[index][0] <= end:
                end = min(points[index][1], end)
            else:
                count += 1
                end = points[index][1]

            index += 1

        return count
