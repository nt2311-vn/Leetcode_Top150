/* The isBadVersion API is defined in the parent class VersionControl.
      bool IsBadVersion(int version); */

public class Solution : VersionControl
{
    public int FirstBadVersion(int n)
    {
        int i = 1,
            j = n;

        while (i <= j)
        {
            int mid = i + (j - i) / 2;
            if (IsBadVersion(mid))
            {
                j = mid - 1;
            }
            else
            {
                i = mid + 1;
            }
        }

        return i;
    }
}
