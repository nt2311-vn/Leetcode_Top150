using System.Collections.Generic;
using System.Linq;

public class Solution
{
    public long FindTheArrayConcVal(int[] nums)
    {
        long result = 0;
        while (nums.Length > 0)
        {
            int i = 0,
                j = nums.Length - 1;

            if (i != j)
            {
                string concat = $"{nums[i]}{nums[j]}";
                int concatVal = int.Parse(concat);
                result += concatVal;
            }
            else
            {
                result += nums[i];
            }

            if (nums.Length > 1)
            {
                nums = nums.Skip(1).Take(nums.Length - 2).ToArray();
            }
            else
            {
                break;
            }
        }

        return result;
    }
}
