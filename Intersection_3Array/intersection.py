from typing import List


class Solution:
    def arraysIntersection(
        self, arr1: List[int], arr2: List[int], arr3: List[int]
    ) -> List[int]:
        intersection: List[int] = []

        for num in arr1:
            if binary_search(num, arr2):
                intersection.append(num)

        index = len(intersection) - 1

        while index >= 0:
            if not binary_search(intersection[index], arr3):
                intersection.remove(intersection[index])

            index -= 1

        return intersection


def binary_search(val: int, target: List[int]) -> bool:
    start, end = 0, len(target) - 1

    while start <= end:
        middle = (start + end) // 2
        if target[middle] < val:
            start = middle + 1
        elif target[middle] > val:
            end = middle - 1
        else:
            return True

    return False
