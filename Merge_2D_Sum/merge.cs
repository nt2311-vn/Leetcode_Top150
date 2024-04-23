using System.Collections.Generic;

public class Solution
{
    public int[][] MergeArrays(int[][] nums1, int[][] nums2)
    {
        List<int[]> result = new List<int[]>();

        int i = 0,
            j = 0;

        while (i < nums1.Length && j < nums2.Length)
        {
            if (nums1[i][0] == nums2[j][0])
            {
                int key = nums1[i][0];
                int sum = nums1[i][1] + nums2[j][1];
                result.Add(new int[] { key, sum });
                i++;
                j++;
            }
            else if (nums1[i][0] > nums2[j][0])
            {
                result.Add(nums2[j]);
                j++;
            }
            else
            {
                result.Add(nums1[i]);
                i++;
            }
        }

        while (i < nums1.Length)
        {
            result.Add(nums1[i]);
            i++;
        }

        while (j < nums2.Length)
        {
            result.Add(nums2[j]);
            j++;
        }

        return result.ToArray();
    }
}
