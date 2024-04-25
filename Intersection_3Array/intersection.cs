using System.Collections.Generic;

public class Solution
{
    public IList<int> ArraysIntersection(int[] arr1, int[] arr2, int[] arr3)
    {
        List<int> result = new List<int>();

        for (int i = 0; i < arr1.Length; i++)
        {
            if (binarySearch(arr1[i], arr2))
            {
                result.Add(arr1[i]);
            }
        }

        for (int i = result.Count - 1; i >= 0; i--)
        {
            if (!binarySearch(result[i], arr3))
            {
                result.RemoveAt(i);
            }
        }

        return result;
    }

    bool binarySearch(int val, int[] target)
    {
        int start = 0,
            end = target.Length - 1;

        while (start <= end)
        {
            int middle = (int)((start + end) / 2);

            if (target[middle] > val)
            {
                end = middle - 1;
            }
            else if (target[middle] < val)
            {
                start = middle + 1;
            }
            else
            {
                return true;
            }
        }

        return false;
    }
}
