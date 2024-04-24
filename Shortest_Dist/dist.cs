using System;
using System.Collections.Generic;

public class Solution
{
    public int[] ShortestToChar(string s, char c)
    {
        List<int> result = new List<int>();

        for (int k = 0; k < s.Length; k++)
        {
            if (s[k] == c)
            {
                result.Add(0);
            }
            else
            {
                int i = k - 1,
                    j = k + 1;

                while (i >= 0 && s[i] != c)
                {
                    i--;
                }

                while (j < s.Length && s[j] != c)
                {
                    j++;
                }

                int minDist = Math.Min(k - i, j - k);

                if (i == -1)
                {
                    minDist = j - k;
                }

                if (j == s.Length)
                {
                    minDist = k - i;
                }

                result.Add(minDist);
            }
        }

        return result.ToArray();
    }
}
