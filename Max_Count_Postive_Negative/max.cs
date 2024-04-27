using System;
using System.Collection.Generic;

public class Solution
{
    public int MaximumCount(int[] nums)
    {
        int lower = binarySearch(nums)[0];
        int upper = binarySearch(nums)[1];

        int negative = 0,
            positive = 0;

        for (int i = 0; i < lower; i++)
        {
            if (nums[i] < 0)
            {
                negative++;
            }
        }

        for (int i = upper; i < nums.Length; i++)
        {
            if (i >= 0 && nums[i] > 0)
            {
                positive++;
            }
        }

        return Math.Max(negative, positive);
    }

    private int[] binarySearch(int[] nums)
    {
        int i = 0,
            j = nums.Length - 1;

        while (i <= j)
        {
            int mid = i + (j - i) / 2;

            if (nums[mid] < 0)
            {
                i = mid + 1;
            }
            else
            {
                j = mid - 1;
            }
        }

        return new int[] { i, j };
    }
}
