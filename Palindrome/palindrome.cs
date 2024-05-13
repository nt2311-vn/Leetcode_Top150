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
    public bool IsPalindrome(ListNode head)
    {
        var slow = head;
        var fast = head;
        ListNode prev = null;

        while (fast != null && fast.next != null)
        {
            fast = fast.next.next;
            ListNode next = slow.next;
            slow.next = prev;
            prev = slow;
            slow = next;
        }

        if (fast != null)
        {
            slow = slow.next;
        }

        while (slow != null)
        {
            if (slow.val != prev.val)
            {
                return fals;
            }

            slow = slow.next;
            prev = prev.next;
        }

        return true;
    }
}
