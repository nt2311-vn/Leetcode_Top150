using System;

public class ListNode
{
    public int val;
    public ListNode next;

    public ListNode(int x)
    {
        val = x;
    }
}

public class Solution
{
    public ListNode GetIntersectionNode(ListNode headA, ListNode headB)
    {
        var a = headA;
        var b = headB;

        while (a != b)
        {
            a = a == null ? headB : a.next;
            b = b == null ? headA : b.next;
        }

        return a;
    }
}
