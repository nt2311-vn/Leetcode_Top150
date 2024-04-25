using System;
using System.Collections.Generic;

public class Solution
{
    public int[] KWeakestRows(int[][] mat, int k)
    {
        List<RowInfo> rowInfo = new List<RowInfo>();

        for (int i = 0; i < mat.Length; i++)
        {
            rowInfo.Add(new RowInfo(i, countSoldier(mat[i])));
        }

        rowInfo.Sort(
            (x, y) =>
            {
                if (x.soldiers == y.soldiers)
                {
                    return x.rowIndex.CompareTo(y.rowIndex);
                }

                return x.soldiers.CompareTo(y.soldiers);
            }
        );

        List<int> result = new List<int>();

        for (int i = 0; i < k; i++)
        {
            result.Add(rowInfo[i].rowIndex);
        }

        return result.ToArray();
    }

    private int countSoldier(int[] row)
    {
        int i = 0,
            j = row.Length - 1;

        while (i <= j)
        {
            int mid = (i + j) / 2;

            if (row[mid] == 1)
            {
                i = mid + 1;
            }
            else
            {
                j = mid - 1;
            }
        }

        return i;
    }

    class RowInfo
    {
        public int rowIndex;
        public int soldiers;

        public RowInfo(int rowIndex, int soldiers)
        {
            this.rowIndex = rowIndex;
            this.soldiers = soldiers;
        }
    }
}
