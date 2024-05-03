public class Solution
{
    public int FixedPoint(int[] arr)
    {
        int i = 0,
            j = arr.Length - 1;

        while (i < j)
        {
            int mid = i + (j - i) / 2;

            if (arr[mid] < mid)
            {
                i = mid + 1;
            }
            else
            {
                j = mid;
            }
        }

        if (arr[i] == i)
        {
            return arr[i];
        }

        return -1;
    }
}
