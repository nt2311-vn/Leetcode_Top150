from typing import List, Tuple


class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        nums = sorted(nums)

        indices: List[int] = []
        low, high = find_bound(nums, target)

        index = low

        while index <= high:
            if nums[index] == target:
                indices.append(index)

            index += 1

        return indices


def find_bound(nums: List[int], target: int) -> Tuple[int, int]:
    i, j = 0, len(nums) - 1

    while i <= j:
        mid = (i + j) // 2

        if nums[mid] > target:
            j = mid - 1
        elif nums[mid] < target:
            i = mid + 1
        else:
            break

    return (i, j)
