from typing import List, Dict


class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        nums1 = sorted(nums1)
        nums2 = sorted(nums2)
        intersect: List[int] = []
        check_table: Dict[int, bool] = {}

        i, j = 0, 0
        while i < len(nums1) and j < len(nums2):
            if nums1[i] == nums2[j]:
                if nums1[i] not in check_table:
                    check_table[nums1[i]] = True
                    intersect.append(nums1[i])

                i += 1
                j += 1

            elif nums1[i] > nums2[j]:
                j += 1
            else:
                i += 1

        return intersect
