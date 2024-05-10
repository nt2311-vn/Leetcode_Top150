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
    public ListNode DeleteNodes(ListNode head, int m, int n)
    {
        ListNode dummyNode = new ListNode(next: head);
        var current = dummyNode;

        while (current != null)
        {
            for (int i = 0; i < m && current != null; i++)
            {
                current = current.next;
            }

            if (current == null)
            {
                return dummyNode.next;
            }

            var nextNode = current;

            for (int i = 0; i < n && nextNode != null; i++)
            {
                nextNode = nextNode.next;
            }

            if (nextNode == null)
            {
                current.next = null;
                return dummyNode.next;
            }

            current.next = nextNode.next;
            current = nextNode;
        }

        return dummyNode.next;
    }
}
