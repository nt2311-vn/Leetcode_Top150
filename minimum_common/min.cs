using System;
using System.Collections.Generic;
using System.Linq;

public class Solution
{
    public int GetCommon(int[] nums1, int[] nums2)
    {
        List<int> common = new List<int>();
        for (int i = 0; i < nums1.Length; i++)
        {
            if (binarySearch(nums2, nums1[i]) != -1)
            {
                common.Add(nums1[i]);
            }
        }

        return common.Count == 0 ? -1 : common.Min();
    }

    int binarySearch(int[] nums, int target)
    {
        int i = 0,
            j = nums.Length - 1;

        while (i <= j)
        {
            int mid = i + (j - i) / 2;

            if (nums[mid] > target)
            {
                j = mid - 1;
            }
            else if (nums[mid] < target)
            {
                i = mid + 1;
            }
            else
            {
                return target;
            }
        }

        return -1;
    }
}
