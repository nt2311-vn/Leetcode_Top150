from typing import List


class Solution:
    def getCommon(self, nums1: List[int], nums2: List[int]) -> int:
        common_list: List[int] = []

        for num in nums1:
            if binary_search(nums2, num):
                common_list.append(num)

        if len(common_list) == 0:
            return -1
        else:
            return min(common_list)


def binary_search(nums: List[int], target: int) -> bool:
    i, j = 0, len(nums) - 1

    while i <= j:
        mid = i + (j - i) // 2

        if nums[mid] < target:
            i = mid + 1
        elif nums[mid] > target:
            j = mid - 1
        else:
            return True

    return False
