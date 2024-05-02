using System;
using System.Collections.Generic;

public class Solution
{
    public int MissingNumber(int[] nums)
    {
        Array.Sort(nums);
        int i = 0,
            j = nums.Length;

        while (i < j)
        {
            int mid = i + (j - i) / 2;
            if (nums[mid] > mid)
            {
                j = mid;
            }
            else
            {
                i = mid + 1;
            }
        }

        return i;
    }
}
