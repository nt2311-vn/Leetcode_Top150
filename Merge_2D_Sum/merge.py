from typing import List


class Solution:
    def mergeArrays(
        self, nums1: List[List[int]], nums2: List[List[int]]
    ) -> List[List[int]]:
        merge: List[List[int]] = []

        i, j = 0, 0

        while i < len(nums1) and j < len(nums2):
            if nums1[i][0] == nums2[j][0]:
                sum = nums1[i][1] + nums2[j][1]
                merge.append([nums1[i][0], sum])
                i, j = i + 1, j + 1
            elif nums1[i][0] > nums2[j][0]:
                merge.append(nums2[j])
                j += 1
            else:
                merge.append(nums1[i])
                i += 1

        while i < len(nums1):
            merge.append(nums1[i])
            i += 1

        while j < len(nums2):
            merge.append(nums2[j])
            j += 1

        return merge
