using System;

public class ListNode
{
    public int val;
    public ListNode next;

    public ListNode(int val = 0, ListNode next = null)
    {
        this.val = val;
        this.next = next;
    }
}

public class Solution
{
    public int GetDecimalValue(ListNode head)
    {
        string binaryStr = "";
        while (head != null)
        {
            binaryStr += head.val.ToString();
            head = head.next;
        }
        return Convert.ToInt32(binaryStr, 2);
    }
}
