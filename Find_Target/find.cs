using System;
using System.Collections.Generic;

public class Solution
{
    public IList<int> TargetIndices(int[] nums, int target)
    {
        List<int> result = new List<int>();
        Array.Sort(nums);

        int i = findLowerBound(nums, target);
        int j = findUpperBoud(nums, target);

        for (int k = i; k <= j; k++)
        {
            if (nums[k] == target)
            {
                result.Add(k);
            }
        }

        return result;
    }

    private int findLowerBound(int[] nums, int target)
    {
        int i = 0,
            j = nums.Length - 1;

        while (i <= j)
        {
            int mid = (i + j) / 2;

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
                break;
            }
        }
        return i;
    }

    private int findUpperBoud(int[] nums, int target)
    {
        int i = 0,
            j = nums.Length - 1;

        while (i <= j)
        {
            int mid = (i + j) / 2;
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
                break;
            }
        }

        return j;
    }
}
