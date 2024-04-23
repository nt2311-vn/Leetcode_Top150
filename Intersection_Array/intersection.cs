using System;
using System.Collections.Generic;

public class Solution
{
    public int[] Intersection(int[] nums1, int[] nums2)
    {
        HashSet<int> result = new HashSet<int>();
        Array.Sort(nums1);
        Array.Sort(nums2);

        int i = 0,
            j = 0;

        while (i < nums1.Length && j < nums2.Length)
        {
            if (nums1[i] == nums2[j])
            {
                result.Add(nums1[i]);
                i++;
                j++;
            }
            else if (nums1[i] > nums2[j])
            {
                j++;
            }
            else
            {
                i++;
            }
        }

        return result.ToArray();
    }
}
