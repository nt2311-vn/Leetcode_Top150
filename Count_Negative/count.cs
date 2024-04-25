using System;
using System.Collections.Generic;

public class Solution
{
    public int CountNegatives(int[][] grid)
    {
        int result = 0;

        for (int i = 0; i < grid.Length; i++)
        {
            result += binarySearch(grid[i]);
        }

        return result;
    }

    int binarySearch(int[] target)
    {
        int i = 0,
            j = target.Length - 1;

        while (i <= j)
        {
            int mid = (int)(i + j) / 2;
            if (target[mid] >= 0)
            {
                i = mid + 1;
            }
            else
            {
                j = mid - 1;
            }
        }

        int count = 0;

        for (int k = i; k < target.Length; k++)
        {
            count++;
        }

        return count;
    }
}
