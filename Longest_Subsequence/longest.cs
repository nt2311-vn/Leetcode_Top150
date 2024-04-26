using System;

public class Solution
{
    public int[] AnswerQueries(int[] nums, int[] queries)
    {
        Array.Sort(nums);
        int[] result = new int[queries.Length];

        for (int i = 0; i < queries.Length; i++)
        {
            int sum = 0,
                subsequenceCount = 0;

            for (int j = 0; j < nums.Length; j++)
            {
                if (sum + nums[j] <= queries[i])
                {
                    sum += nums[j];
                    subsequenceCount++;
                }
                else
                {
                    break;
                }
            }
            result[i] = subsequenceCount;
        }

        return result;
    }
}
