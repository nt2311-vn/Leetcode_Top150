using System;
using System.Collections.Generic;
using System.Linq;

public class Solution
{
    public int[] SortedSquares(int[] nums)
    {
        int[] result = new int[nums.Length];

        int i = 0,
            j = nums.Length - 1,
            index = nums.Length - 1;

        while (i <= j)
        {
            int iSquare = nums[i] * nums[i];
            int jSquare = nums[j] * nums[j];
            if (iSquare > jSquare)
            {
                result[index] = iSquare;
                i++;
            }
            else
            {
                result[index] = jSquare;
                j--;
            }

            index--;
        }

        return result;
    }
}
