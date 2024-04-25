from typing import List


class Solution:
    def countNegatives(self, grid: List[List[int]]) -> int:
        result = 0

        for arr in grid:
            result += binary_search(arr)

        return result


def binary_search(target: List[int]) -> int:
    i, j = 0, len(target) - 1

    while i <= j:
        mid = (i + j) // 2

        if target[mid] >= 0:
            i = mid + 1
        else:
            j = mid - 1

    return len(target[i:])
