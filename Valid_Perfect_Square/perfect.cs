public class Solution
{
    public bool IsPerfectSquare(int num)
    {
        if (num < 4)
        {
            return num == 1 ? true : false;
        }

        int i = 0,
            j = num / 2;

        while (i <= j)
        {
            int mid = i + (j - i) / 2;

            if (mid * mid == num)
            {
                return true;
            }
            else if (mid < num / mid)
            {
                i = mid + 1;
            }
            else
            {
                j = mid - 1;
            }
        }

        return false;
    }
}
