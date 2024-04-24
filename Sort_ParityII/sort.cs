using System.Collections.Generic;

public class Solution
{
    public int[] SortArrayByParityII(int[] nums)
    {
        int[] result = new int[nums.Length];
        int i = 0,
            j = 1;

        for (int k = 0; k < nums.Length; k++)
        {
            if (nums[k] % 2 == 0)
            {
                result[i] = nums[k];
                i += 2;
            }
            else
            {
                result[j] = nums[k];
                j += 2;
            }
        }

        return result;
    }
}
